package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/liip/sheriff"
	"github.com/spf13/viper"
)

const (
	// defines environment states

	// EnvStateRunning means that ALL instances for an env are in "running" state
	EnvStateRunning = "running"
	// EnvStateStopped means that ALL instances for an env are in "stopped" state
	EnvStateStopped = "stopped"
	// EnvStateMixed means that instances for an env are in EITHER "stopped" or "running" state
	EnvStateMixed = "mixed"
	// EnvStateChanging means that AT LEAST ONE instance for an env is NOT in a "running" state or "stopped" state
	EnvStateChanging = "changing"
	// ASGLabel is used to identify an ASG instance
	ASGLabel = "ASG"
)

var (
	// global aws clients (based on regions)
	awsClients map[string]*ec2.Client
	// global autoscaling clients (based on regions)
	awsASGClients map[string]*autoscaling.Client
	// global cached env list
	cachedTable envList
	// lock to prevent concurrent refreshes
	cachedTableLock sync.Mutex
	// aws regions
	awsRegions []string
	// aws tags
	requiredTagKey, requiredTagValue, environmentTagKey string
	// safety, will refuse to shutdown if more than this amount of instances is requested
	maxInstancesToShutdown int
	// ignore these instance types
	instanceTypeIgnore []string
	// ignore these environment names
	envNameIgnore []string
	// aws api poll interval
	pollInterval time.Duration
	// bills accrued per env
	billsAccruedMap = map[string]float64{}
	// bills saved per env
	billsSavedMap = map[string]float64{}
	// total bills accrued since aws-power-toggle server start
	totalBillsAccrued float64
	// total bills saved since aws-power-toggle server start
	totalBillsSaved float64
	// instance IDs the server has actually toggled off
	toggledOffInstanceIds = map[string]bool{}
	// lock to prevent concurrent access of the above map
	toggledOffInstanceIdsLock sync.RWMutex
	// last time aws api was accessed
	lastRefreshedTimeUnixNano int64
	// MockEnabled enable mocking of API calls to aws for development purposes
	mockEnabled bool
	// experimentalEnabled enable experimental features. Currently include billing stats
	experimentalEnabled bool
	// enable support for interacting with ASGs (Auto Scaling Groups)
	asgEnabled bool
)

type virtualMachine struct {
	// ID unique to this application
	ID string `json:"id" groups:"summary,details"`

	// these values are straight from aws api
	InstanceID   string `json:"instance_id" groups:"summary,details"`
	InstanceType string `json:"instance_type" groups:"summary,details"`
	Name         string `json:"name" groups:"summary,details"`
	State        string `json:"state" groups:"summary,details"`
	Environment  string `json:"environment" groups:"summary,details"`
	Region       string `json:"region" groups:"summary,details"`

	// these values are mapped from another source for aws
	VCPU          int     `json:"vcpu" groups:"summary,details"`
	MemoryGB      float32 `json:"memory_gb" groups:"summary,details"`
	PricingHourly float64 `json:"pricing" groups:"summary,details"`

	// ASG values
	IsASG            bool  `json:"is_asg groups:"summary,details""`
	ASGInstanceCount int   `json:"asg_instance_count" groups:"summary,details"`
	MinSize          int64 `json:"min_size" groups:"summary,details"`
	MaxSize          int64 `json:"max_size" groups:"summary,details"`
	DesiredCapacity  int64 `json:"desired_capacity" groups:"summary,details"`
}

type environment struct {
	// ID unique to this application
	ID        string           `json:"id" groups:"summary,details"`
	Provider  string           `json:"provider" groups:"summary,details"`
	Region    string           `json:"region" groups:"summary,details"`
	Name      string           `json:"name" groups:"summary,details"`
	Instances []virtualMachine `json:"instances" groups:"details"`

	// these values are calculated based on list of Instances
	RunningInstances int     `json:"running_instances" groups:"summary,details"`
	StoppedInstances int     `json:"stopped_instances" groups:"summary,details"`
	TotalInstances   int     `json:"total_instances" groups:"summary,details"`
	TotalVCPU        int     `json:"total_vcpu" groups:"summary,details"`
	TotalMemoryGB    float32 `json:"total_memory_gb" groups:"summary,details"`
	State            string  `json:"state" groups:"summary,details"`
	BillsAccrued     string  `json:"bills_accrued,omitempty" groups:"summary,details"`
	BillsSaved       string  `json:"bills_saved,omitempty" groups:"summary,details"`
}

// for global cached table
type envList []environment

// updateEnvDetails
// determines details like: State, TotalVCPU, TotalMemoryGB
func updateEnvDetails() {
	for i, env := range cachedTable {
		// these are not fully used yet (currently only for api response)
		// in future we may support other providers and multiple regions
		cachedTable[i].Provider = "aws"

		// compute a unique identifier for this environment
		cachedTable[i].ID = ComputeID(
			cachedTable[i].Provider,
			cachedTable[i].Region,
			env.Name,
		)
		// add bills accrued and bills saved to env details
		if experimentalEnabled {
			if envbillAccrued, exists := billsAccruedMap[cachedTable[i].ID]; exists {
				cachedTable[i].BillsAccrued = fmt.Sprintf("%.02f", envbillAccrued)
			}
			if envbillSaved, exists := billsSavedMap[cachedTable[i].ID]; exists {
				cachedTable[i].BillsSaved = fmt.Sprintf("%.02f", envbillSaved)
			}
		}

		// vm total count
		cachedTable[i].TotalInstances = len(env.Instances)
		// reset counts
		cachedTable[i].RunningInstances = 0
		cachedTable[i].StoppedInstances = 0
		cachedTable[i].TotalVCPU = 0
		cachedTable[i].TotalMemoryGB = 0

		// determine counts
		for c, instance := range env.Instances {
			// compute a unique identifier for this instance
			//   InstanceID is already unique, but this will make ids consistent
			//   in case we add other cloud providers
			cachedTable[i].Instances[c].ID = ComputeID(
				cachedTable[i].Provider,
				instance.Region,
				instance.InstanceID,
			)

			// update cpu and memory counts
			cachedTable[i].TotalVCPU += instance.VCPU
			cachedTable[i].TotalMemoryGB += instance.MemoryGB
			// update vm states
			switch instance.State {
			case "running":
				cachedTable[i].RunningInstances++
			case "stopped":
				cachedTable[i].StoppedInstances++
			}
		}

		// determine environment state
		switch {
		case cachedTable[i].TotalInstances == cachedTable[i].RunningInstances:
			cachedTable[i].State = EnvStateRunning
		case cachedTable[i].TotalInstances == cachedTable[i].StoppedInstances:
			cachedTable[i].State = EnvStateStopped
		case cachedTable[i].TotalInstances == (cachedTable[i].RunningInstances + cachedTable[i].StoppedInstances):
			cachedTable[i].State = EnvStateMixed
		default:
			cachedTable[i].State = EnvStateChanging
		}
	}
}

// calculateEnvBills calculate bills accrued / saved since the last aws poll
// return a map of env IDs with their respective bills accrued/saved
func calculateEnvBills() {
	// acquire and release lock on instance id map only once here, to avoid doing it for every map read
	toggledOffInstanceIdsLock.RLock()
	defer toggledOffInstanceIdsLock.RUnlock()

	newRefreshedTimeUnixNano := time.Now().UnixNano()
	elapsedTimeInHour := (float64(newRefreshedTimeUnixNano) - float64(lastRefreshedTimeUnixNano)) * float64(time.Nanosecond) / float64(time.Hour)
	lastRefreshedTimeUnixNano = newRefreshedTimeUnixNano
	for _, env := range cachedTable {
		var envBillsAccrued float64
		var envBillsSaved float64
		for _, instance := range env.Instances {
			// calculate bills as if state of instances are unchanged since the last poll. It's the best we can do for now (or so I believe)
			instanceBill := instance.PricingHourly * elapsedTimeInHour
			switch instance.State {
			case "running":
				envBillsAccrued += instanceBill
			case "stopped":
				// before claiming any responsibilities, need to find out whether the instance was actually stopped by aws-power-toggle :)
				if toggledOffInstanceIds[instance.InstanceID] {
					envBillsSaved += instanceBill
				}
			default:
			}
		}
		billsAccruedMap[env.ID] = billsAccruedMap[env.ID] + envBillsAccrued
		billsSavedMap[env.ID] = billsSavedMap[env.ID] + envBillsSaved
		totalBillsAccrued += envBillsAccrued
		totalBillsSaved += envBillsSaved
	}
	return
}

// checks if an instance should be included based on instance type
// true if its OK, false to ignore
func checkInstanceType(instanceType string) (ok bool) {
	ok = true
	for _, ignoredType := range instanceTypeIgnore {
		if ignoredType == instanceType {
			ok = false
			break
		}
	}
	return
}

// checks if an instance should be included based on environment name
// also ensures that the env name is not empty
// true if its OK, false to ignore
func validateEnvName(envName string) (ok bool) {
	ok = true
	if envName == "" {
		ok = false
		return
	}
	for _, ignoredEnvName := range envNameIgnore {
		if ignoredEnvName == envName {
			ok = false
			break
		}
	}
	return
}

// adds and instance to the global cachedTable
func addInstance(instance *virtualMachine) {
	// check if we should ignore instance based on:
	//  - instance is not part of an ASG
	//  - instance type is not on ignore list or in terminated state
	if !instance.IsASG && (!checkInstanceType(instance.InstanceType) || instance.State == "terminated") {
		log.Debugf("instance is being ignored: name='%s' [%s](%s)\n", instance.Name, instance.InstanceType, instance.State)
		return
	}
	envExists := false
	for i, env := range cachedTable {
		if env.Name == instance.Environment {
			envExists = true
			cachedTable[i].Instances = append(cachedTable[i].Instances, *instance)
		}
	}
	if !envExists {
		ec2env := environment{
			Name:      instance.Environment,
			Instances: []virtualMachine{*instance},
			Region:    instance.Region,
		}
		cachedTable = append(cachedTable, ec2env)
	}
}

// returns a list of discovered ASGs. Each instance returned will be a single ASG
// with a cumulative total given for instance vCPU and memory
func pollForASG() (instances []virtualMachine, err error) {
	param := &autoscaling.DescribeAutoScalingGroupsInput{}
	for region, awsASvcClient := range awsASGClients {
		pollASGStartTime := time.Now()
		req := awsASvcClient.DescribeAutoScalingGroupsRequest(param)
		resp, respErr := req.Send(context.Background())
		if respErr != nil {
			log.Errorf("failed to describe AutoScalingGroups, %s, %v", region, respErr)
			err = respErr
			return
		}
		for _, asg := range resp.AutoScalingGroups {
			instanceObj := virtualMachine{
				IsASG: true,
				// by default we use the asg name for the "instance" name.
				// We will ignore the Name tag
				Name:             *asg.AutoScalingGroupName,
				InstanceID:       ASGLabel,
				InstanceType:     ASGLabel,
				Region:           region,
				ASGInstanceCount: len(asg.Instances),
				MinSize:          *asg.MinSize,
				MaxSize:          *asg.MaxSize,
				DesiredCapacity:  *asg.DesiredCapacity,
			}

			isValidASG := false
			for _, tag := range asg.Tags {
				if *tag.Key == "power-toggle-enabled" && *tag.Value == "true" {
					isValidASG = true
					// gather some additional information about this ASG
					if len(asg.Instances) > 0 && *asg.DesiredCapacity > 0 {
						instanceObj.State = "running"
						for _, i := range asg.Instances {
							// We sum the memory and vcpu of all the instances in an ASG (they appear as a single entry)
							if details, found := getInstanceTypeDetails(*i.InstanceType); found {
								instanceObj.MemoryGB += details.MemoryGB
								instanceObj.VCPU += details.VCPU
								if pricingStr, ok := details.PricingHourlyByRegion[region]; ok {
									pricing, errPrice := strconv.ParseFloat(pricingStr, 64)
									if errPrice != nil {
										log.Errorf("failed to parse pricing info to float: %s", pricingStr)
									}
									instanceObj.PricingHourly = pricing
								}
							}
						}
					} else {
						instanceObj.State = "stopped"
					}
				}
				if *tag.Key == environmentTagKey && *tag.Value != "" {
					instanceObj.Environment = *tag.Value
				}
			}
			if isValidASG && validateEnvName(instanceObj.Environment) {
				// if the ASG matches tags we add it like if it was a EC2.
				instances = append(instances, instanceObj)
			}
		}
		elapsed := time.Since(pollASGStartTime)
		log.Debugf("polling for ASGs in region %s took %s", region, elapsed)
	}
	return
}

// returns a list of discovered EC2 instances.
func pollForEC2() (instances []virtualMachine, err error) {
	params := &ec2.DescribeInstancesInput{
		Filters: []ec2.Filter{
			{
				Name: aws.String("tag:" + requiredTagKey),
				Values: []string{
					requiredTagValue,
				},
			},
		},
	}

	for region, awsSvcClient := range awsClients {
		pollEC2StartTime := time.Now()
		req := awsSvcClient.DescribeInstancesRequest(params)
		resp, respErr := req.Send(context.Background())
		if respErr != nil {
			log.Errorf("failed to describe instances, %s, %v", region, respErr)
			err = respErr
			return
		}

		for _, reservation := range resp.Reservations {
			for _, instance := range reservation.Instances {
				instanceObj := virtualMachine{
					InstanceID: *instance.InstanceId, State: string(instance.State.Name),
					InstanceType: string(instance.InstanceType),
					Region:       region,
				}
				// populate info from tags
				isASG := false
				for _, tag := range instance.Tags {
					if *tag.Key == environmentTagKey && *tag.Value != "" {
						instanceObj.Environment = *tag.Value
					}
					if *tag.Key == "Name" {
						instanceObj.Name = *tag.Value
					}
					if *tag.Key == "aws:autoscaling:groupName" {
						isASG = true
					}
				}
				// if true Instance is part of ASG. bypass this instance
				if isASG {
					continue // goto next instance
				}
				// determine instance cpu and memory
				if details, found := getInstanceTypeDetails(instanceObj.InstanceType); found {
					instanceObj.MemoryGB = details.MemoryGB
					instanceObj.VCPU = details.VCPU
					if pricingstr, ok := details.PricingHourlyByRegion[region]; ok {
						pricing, err := strconv.ParseFloat(pricingstr, 64)
						if err != nil {
							log.Errorf("failed to parse pricing info to float: %s", pricingstr)
						}
						instanceObj.PricingHourly = pricing
					}
				}
				if validateEnvName(instanceObj.Environment) {
					instances = append(instances, instanceObj)
				}
			}

		}
		elapsed := time.Since(pollEC2StartTime)
		log.Debugf("polling for EC2s in region %s took %s", region, elapsed)
	}
	return
}

// polls aws for updates to cachedTable
func refreshTable() (err error) {
	cachedTableLock.Lock()
	defer cachedTableLock.Unlock()

	// use the mock function if enabled
	if mockEnabled {
		return mockRefreshTable()
	}

	// calculate billing information before old table is ditched
	if experimentalEnabled {
		calculateEnvBills()
	}

	// used to calculate the time it took to poll aws
	pollStartTime := time.Now()

	// keeps track of everything we discovered during this poll
	var discoveredInstances []virtualMachine

	// discover ASGs (Auto Scaling Groups) if enabled
	if asgEnabled {
		if discoveredASGs, err := pollForASG(); err == nil {
			discoveredInstances = append(discoveredInstances, discoveredASGs...)
		} else {
			return fmt.Errorf("error polling ASG: %v", err)
		}
	}

	// discover EC2 Instances
	if discoveredEC2Instances, err := pollForEC2(); err == nil {
		discoveredInstances = append(discoveredInstances, discoveredEC2Instances...)
	} else {
		return fmt.Errorf("error polling EC2: %v", err)
	}

	// polling was successful, now we rebuild the cache
	cachedTable = cachedTable[:0]
	for _, discoveredInstance := range discoveredInstances {
		addInstance(&discoveredInstance)
	}
	updateEnvDetails()

	elapsed := time.Since(pollStartTime)
	log.Debugf("total polling time took %s; valid environment(s) in cache: %d", elapsed, len(cachedTable))
	return
}

// get instance ids for an environment with a specific state
// this is used for power up/down commands against aws API
func getInstanceIDs(envID, state string) (instanceIds []string) {
	for _, env := range cachedTable {
		if env.ID == envID {
			for _, instance := range env.Instances {
				if !instance.IsASG && instance.State == state {
					instanceIds = append(instanceIds, instance.InstanceID)
				}
			}
		}
	}
	return
}

// get ASG name for an environment with a specific state
// this is used for power up/down commands against aws API
func getASGs(envID, state string) (asgNames []string, instanceCount int) {
	for _, env := range cachedTable {
		if env.ID == envID {
			for _, instance := range env.Instances {
				if instance.IsASG && instance.State == state {
					asgNames = append(asgNames, instance.Name)
					instanceCount += instance.ASGInstanceCount
				}
			}
		}
	}
	return
}

// toggleInstances can start or stop a list of instances
func toggleInstances(instanceIDs []string, desiredState string, awsClient *ec2.Client) (response []byte, err error) {
	if len(instanceIDs) < 1 {
		err = fmt.Errorf("no instanceIDs have been provided")
		return
	}

	// supported states are: start, stop
	switch desiredState {
	case "start":
		input := &ec2.StartInstancesInput{
			InstanceIds: instanceIDs,
			DryRun:      aws.Bool(false),
		}

		req := awsClient.StartInstancesRequest(input)
		awsResponse, reqErr := req.Send(context.Background())
		response, _ = json.MarshalIndent(awsResponse, "", "  ")
		err = reqErr
		if experimentalEnabled && err == nil {
			// BILLING: update toggled off instances map
			deleteToggledOffInstanceIDs(instanceIDs)
		}
		return

	case "stop":
		input := &ec2.StopInstancesInput{
			InstanceIds: instanceIDs,
			DryRun:      aws.Bool(false),
		}

		req := awsClient.StopInstancesRequest(input)
		awsResponse, reqErr := req.Send(context.Background())
		response, _ = json.MarshalIndent(awsResponse, "", "  ")
		err = reqErr
		if experimentalEnabled && err == nil {
			// BILLING: update toggled off instances map
			putToggledOffInstanceIDs(instanceIDs)
		}
		return

	default:
		err = fmt.Errorf("unsupported desiredState specified")
		return
	}
}

// toggleInstances can start or stop a list of ASGs
func toggleASGs(asgNames []string, desiredState string, awsASGClient *autoscaling.Client) (response []byte, err error) {
	if len(asgNames) < 1 {
		err = fmt.Errorf("no ASG names have been provided")
		return
	}

	// supported states are: start, stop
	switch desiredState {
	case "start":
		for _, asg := range asgNames {
			// Must: DesiredCapacity >= MinSize , need to set both
			// At start setting ASG to 1 as we haven't cached the original value.
			input := &autoscaling.UpdateAutoScalingGroupInput{
				AutoScalingGroupName: aws.String(asg),
				DesiredCapacity:      aws.Int64(1),
				MinSize:              aws.Int64(1),
			}
			req := awsASGClient.UpdateAutoScalingGroupRequest(input)
			awsResponse, reqErr := req.Send(context.Background())
			response, _ = json.MarshalIndent(awsResponse, "", "  ")
			err = reqErr
			if experimentalEnabled && err == nil {
				// BILLING: update toggled off instances map
				putToggledOffInstanceIDs(asgNames)
			}
		}
		return
	case "stop":
		for _, asg := range asgNames {
			// Must: DesiredCapacity >= MinSize , need to set both
			input := &autoscaling.UpdateAutoScalingGroupInput{
				AutoScalingGroupName: aws.String(asg),
				DesiredCapacity:      aws.Int64(0),
				MinSize:              aws.Int64(0),
			}
			req := awsASGClient.UpdateAutoScalingGroupRequest(input)
			awsResponse, reqErr := req.Send(context.Background())
			response, _ = json.MarshalIndent(awsResponse, "", "  ")
			err = reqErr
			if experimentalEnabled && err == nil {
				// BILLING: update toggled off instances map
				putToggledOffInstanceIDs(asgNames)
			}
		}
		return
	default:
		err = fmt.Errorf("unsupported desiredState specified")
		return
	}
}

func putToggledOffInstanceIDs(instanceIDs []string) {
	toggledOffInstanceIdsLock.Lock()
	for _, instanceID := range instanceIDs {
		toggledOffInstanceIds[instanceID] = true
	}
	toggledOffInstanceIdsLock.Unlock()
}

func deleteToggledOffInstanceIDs(instanceIDs []string) {
	toggledOffInstanceIdsLock.Lock()
	for _, instanceID := range instanceIDs {
		delete(toggledOffInstanceIds, instanceID)
	}
	toggledOffInstanceIdsLock.Unlock()
}

// shuts down an env
func shutdownEnv(envID string) (response []byte, err error) {
	// use the mock function if enabled
	if mockEnabled {
		return mockShutdownEnv(envID)
	}

	// get env details
	env, found := getEnvironmentByID(envID)
	if !found {
		log.Errorf("env ID %s was not found in the cache", envID)
		return
	}

	// get ASGs for this environment
	asgNames, asgInstanceCount := getASGs(envID, "running")
	// get instance IDs for this environment
	instanceIds := getInstanceIDs(envID, "running")
	// calculate the total instances running in this environment
	totalInstanceCount := len(instanceIds) + asgInstanceCount
	if totalInstanceCount > maxInstancesToShutdown {
		err = fmt.Errorf("SAFETY: env %s [%s] has too many associated instances to shutdown %d", env.Name, envID, totalInstanceCount)
		log.Debugf("SAFETY: instances: %v", instanceIds)
		return
	}

	// shutdown non-ASG EC2 instances
	var errInstance error
	if len(instanceIds) > 0 {
		if _, errInstance = toggleInstances(instanceIds, "stop", getEnvironmentAwsClient(envID)); errInstance != nil {
			log.Errorf("error trying to stop instances for env %s [%s]: %v", env.Name, envID, errInstance)
		}
	}

	// shutdown ASGs
	var errASG error
	if len(asgNames) > 0 && true {
		if _, errASG = toggleASGs(asgNames, "stop", getEnvironmentAwsASGClient(envID)); errASG != nil {
			log.Errorf("error trying to stop ASGs for env %s [%s]: %v", env.Name, envID, errASG)
		}
	}

	// determine if there's any errors
	if errInstance == nil && errASG == nil {
		log.Infof("successfully stopped env %s [%s]", env.Name, envID)
		slackSendMessage(
			fmt.Sprintf(
				"*STOPPING* environment *`%s`* in region _%s_ --> *%v instance(s)* totaling *%v vCPU(s)* & *%vGB* memory",
				env.Name,
				env.Region,
				env.TotalInstances,
				env.TotalVCPU,
				env.TotalMemoryGB,
			),
		)
	} else {
		err = fmt.Errorf("%v %v", errInstance, errASG)
		slackSendMessage(
			fmt.Sprintf(
				"*ERROR STOPPING* environment *`%s`* in region _%s_ --> `%v`",
				env.Name,
				env.Region,
				err,
			),
		)
	}
	return
}

// starts up an env
func startupEnv(envID string) (response []byte, err error) {
	// use the mock function if enabled
	if mockEnabled {
		return mockStartupEnv(envID)
	}

	// get env details
	env, found := getEnvironmentByID(envID)
	if !found {
		log.Errorf("env ID %s was not found in the cache", envID)
		return
	}

	// get ASGs for this environment
	asgNames, _ := getASGs(envID, "stopped")
	// get instance IDs for this environment
	instanceIds := getInstanceIDs(envID, "stopped")

	// start non-ASG EC2 instances
	var errInstance error
	if len(instanceIds) > 0 {
		if _, errInstance = toggleInstances(instanceIds, "start", getEnvironmentAwsClient(envID)); errInstance != nil {
			log.Errorf("error trying to start instances for env %s [%s]: %v", env.Name, envID, errInstance)
		}
	}

	// start ASGs
	var errASG error
	if len(asgNames) > 0 && true {
		if _, errASG = toggleASGs(asgNames, "start", getEnvironmentAwsASGClient(envID)); errASG != nil {
			log.Errorf("error trying to start ASGs for env %s [%s]: %v", env.Name, envID, errASG)
		}
	}

	// determine if there's any errors
	if errInstance == nil && errASG == nil {
		log.Infof("successfully started env %s [%s]", env.Name, envID)
		slackSendMessage(
			fmt.Sprintf(
				"*STARTING* environment *`%s`* in region _%s_ --> *%v instance(s)* totaling *%v vCPU(s)* & *%vGB* memory",
				env.Name,
				env.Region,
				env.TotalInstances,
				env.TotalVCPU,
				env.TotalMemoryGB,
			),
		)
	} else {
		err = fmt.Errorf("%v %v", errInstance, errASG)
		slackSendMessage(
			fmt.Sprintf(
				"*ERROR STARTING* environment *`%s`* in region _%s_ --> `%v`",
				env.Name,
				env.Region,
				err,
			),
		)
	}
	return
}

// starts up an instance based on internal id (not aws instance id)
func toggleInstance(id, desiredState string) (response []byte, err error) {
	// use the mock function if enabled
	if mockEnabled {
		return mockToggleInstance(id, desiredState)
	}

	// validate desiredState
	if desiredState != "start" && desiredState != "stop" {
		err = fmt.Errorf("invalid desired state: %s", desiredState)
		return
	}
	// get the AWS instance id
	awsInstanceID := getAWSInstanceID(id)
	// Here is where we diff between EC2 and ASG, don't name ASG beginning with "i-"
	if strings.HasPrefix(awsInstanceID, "i-") {
		ec2Client := getInstanceAwsClient(id)
		if ec2Client == nil {
			err = fmt.Errorf("could not retrieve an ec2 client for id %s (%s)", id, awsInstanceID)
			return
		}
		response, err = toggleInstances([]string{awsInstanceID}, desiredState, ec2Client)
		if err != nil {
			log.Errorf("error trying to %s instance %s: %v", desiredState, id, err)
		} else {
			log.Infof("successfully toggled instance state (%s): %s", desiredState, id)
		}
		// if the prefix doesn't start with "i-" and it's not empty, it should be an ASG
	} else if awsInstanceID != "" {
		asgClient := getAwsASGClient(awsInstanceID)
		if asgClient == nil {
			err = fmt.Errorf("could not retrieve client for an ASG named: %s", awsInstanceID)
			return
		}
		response, err = toggleASGs([]string{awsInstanceID}, desiredState, asgClient)
		if err != nil {
			log.Errorf("error trying to %s instance %s: %v", desiredState, id, err)
		} else {
			log.Debugf("successfully toggled ASG (%s): %s %s", desiredState, id, awsInstanceID)
		}
	} else {
		err = fmt.Errorf("no mapping found between internal id (%s) and an aws instance id", id)
	}
	return
}

// returns a single environment by id
func getEnvironmentByID(envID string) (environment, bool) {
	for _, env := range cachedTable {
		if env.ID == envID {
			return env, true
		}
	}
	return environment{}, false
}

// returns awsClient for the specific environment ID
func getEnvironmentAwsClient(envID string) *ec2.Client {
	for _, env := range cachedTable {
		if env.ID == envID {
			return awsClients[env.Region]
		}
	}
	return nil
}

// returns awsASGClient for the specific environment ID
func getEnvironmentAwsASGClient(envID string) *autoscaling.Client {
	for _, env := range cachedTable {
		if env.ID == envID {
			return awsASGClients[env.Region]
		}
	}
	return nil
}

// returns awsClient for the specific environment ID
func getInstanceAwsClient(instanceID string) *ec2.Client {
	for _, env := range cachedTable {
		for _, instance := range env.Instances {
			if instance.ID == instanceID {
				return awsClients[instance.Region]
			}
		}
	}
	return nil
}

// returns awsASGClient for the specific instanceID
func getAwsASGClient(asgName string) *autoscaling.Client {
	for _, env := range cachedTable {
		for _, instance := range env.Instances {
			if instance.IsASG && instance.Name == asgName {
				return awsASGClients[instance.Region]
			}
		}
	}
	return nil
}

// given an aws-power-toggle id, it will return the actual aws instance id OR the name of the ASG
func getAWSInstanceID(id string) (awsInstanceID string) {
	for _, env := range cachedTable {
		for _, instance := range env.Instances {
			if instance.ID == id {
				if instance.IsASG {
					awsInstanceID = instance.Name
				} else {
					awsInstanceID = instance.InstanceID
				}
				break
			}
		}
	}
	return
}

// getMarshaledResponse will filter out fields from the struct based on predefined groups
func getMarshaledResponse(data interface{}, groups ...string) (response []byte, err error) {
	// filter out the specified group(s)
	if sMarshal, sErr := sheriff.Marshal(&sheriff.Options{Groups: groups}, data); sErr != nil {
		log.Errorf("error parsing json (sheriff): %v", sErr)
		err = sErr
	} else {
		response, err = json.Marshal(sMarshal)
	}
	return
}

// StartPoller is an infinite loop which periodically polls AWS to refresh the cache
func StartPoller() {
	// build the initial cache
	refreshTable()

	pollInterval = time.Minute * time.Duration(viper.GetInt("aws.polling_interval"))
	log.Infof("start polling with interval %v", pollInterval)

	t := time.Tick(pollInterval)
	// start polling forever...
	for {
		select {
		// interval reached
		case <-t:
			refreshTable()
		}
	}
}
