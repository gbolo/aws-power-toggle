package backend

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/liip/sheriff"
	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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
)

var (

	// global aws clients (based on regions)
	awsClients map[string]*ec2.EC2
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
	MockEnabled bool
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
	BillsAccrued     string  `json:"bills_accrued" groups:"summary,details"`
	BillsSaved       string  `json:"bills_saved" groups:"summary,details"`
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
		if envbillAccrued, exists := billsAccruedMap[cachedTable[i].ID]; exists {
			cachedTable[i].BillsAccrued = fmt.Sprintf("%.02f", envbillAccrued)
		}
		if envbillSaved, exists := billsSavedMap[cachedTable[i].ID]; exists {
			cachedTable[i].BillsSaved = fmt.Sprintf("%.02f", envbillSaved)
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

// adds and instance to cachedTable
func addInstance(instance *virtualMachine) {
	// check if we should ignore instance based on:
	//  - configured ignored instance types
	//  - instance state is "terminated"
	if !checkInstanceType(instance.InstanceType) || instance.State == "terminated" {
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

// polls aws for updates to cachedTable
func refreshTable() (err error) {
	cachedTableLock.Lock()
	defer cachedTableLock.Unlock()

	// use the mock function if enabled
	if MockEnabled {
		return mockRefreshTable()
	}

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

	// calculate billing information before old table is ditched
	calculateEnvBills()

	for region, awsSvcClient := range awsClients {
		req := awsSvcClient.DescribeInstancesRequest(params)
		resp, respErr := req.Send()
		if respErr != nil {
			log.Errorf("failed to describe instances, %s, %v", region, respErr)
			err = respErr
			return
		}
		log.Infof("aws poll was successful, clearing old cached table for region: %s", region)
		var newcachedTable envList
		for _, env := range cachedTable {
			if env.Region != region {
				newcachedTable = append(newcachedTable, env)
			}
		}
		cachedTable = newcachedTable

		for _, reservation := range resp.Reservations {
			for _, instance := range reservation.Instances {
				instanceObj := virtualMachine{
					InstanceID: *instance.InstanceId, State: string(instance.State.Name),
					InstanceType: string(instance.InstanceType),
					Region:       region,
				}
				// populate info from tags
				for _, tag := range instance.Tags {
					if *tag.Key == environmentTagKey && *tag.Value != "" {
						instanceObj.Environment = *tag.Value
					}
					if *tag.Key == "Name" {
						instanceObj.Name = *tag.Value
					}
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
					addInstance(&instanceObj)
				}
			}
		}
		updateEnvDetails()
		log.Debugf("valid environment(s) in cache: %d", len(cachedTable))
	}

	return
}

// get instance ids for an environment with a specific state
// this is used for power up/down commands against aws API
func getInstanceIDs(envID, state string) (instanceIds []string) {
	for _, env := range cachedTable {
		if env.ID == envID {
			for _, instance := range env.Instances {
				if instance.State == state {
					instanceIds = append(instanceIds, instance.InstanceID)
				}
			}
		}
	}
	return
}

// toggleInstances can start or stop a list of instances
func toggleInstances(instanceIDs []string, desiredState string, awsClient *ec2.EC2) (response []byte, err error) {
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
		awsResponse, reqErr := req.Send()
		response, _ = json.MarshalIndent(awsResponse, "", "  ")
		err = reqErr
		if err == nil {
			// BILLING: update toggled off instances map
			toggledOffInstanceIdsLock.Lock()
			for _, instanceID := range instanceIDs {
				delete(toggledOffInstanceIds, instanceID)
			}
			toggledOffInstanceIdsLock.Unlock()
		}
		return

	case "stop":
		input := &ec2.StopInstancesInput{
			InstanceIds: instanceIDs,
			DryRun:      aws.Bool(false),
		}

		req := awsClient.StopInstancesRequest(input)
		awsResponse, reqErr := req.Send()
		response, _ = json.MarshalIndent(awsResponse, "", "  ")
		err = reqErr
		if err == nil {
			// BILLING: update toggled off instances map
			toggledOffInstanceIdsLock.Lock()
			for _, instanceID := range instanceIDs {
				toggledOffInstanceIds[instanceID] = true
			}
			toggledOffInstanceIdsLock.Unlock()
		}
		return

	default:
		err = fmt.Errorf("unsupported desiredState specified")
		return
	}
}

// shuts down an env
func shutdownEnv(envID string) (response []byte, err error) {
	// use the mock function if enabled
	if MockEnabled {
		return mockShutdownEnv(envID)
	}

	instanceIds := getInstanceIDs(envID, "running")
	if len(instanceIds) > maxInstancesToShutdown {
		err = fmt.Errorf("SAFETY: env [%s] has too many associated instances to shutdown %d", envID, len(instanceIds))
		log.Debugf("SAFETY: instances: %v", instanceIds)
	} else if len(instanceIds) > 0 {
		response, err = toggleInstances(instanceIds, "stop", getEnvironmentAwsClient(envID))
		if err != nil {
			log.Errorf("error trying to stop env %s: %v", envID, err)
		} else {
			log.Infof("successfully stopped env %s", envID)
		}
	} else {
		err = fmt.Errorf("env [%s] has no associated instances", envID)
		log.Errorf("env [%s] has no associated instances", envID)
	}
	return
}

// starts up an env
func startupEnv(envID string) (response []byte, err error) {
	// use the mock function if enabled
	if MockEnabled {
		return mockStartupEnv(envID)
	}

	instanceIds := getInstanceIDs(envID, "stopped")
	if len(instanceIds) > 0 {
		response, err = toggleInstances(instanceIds, "start", getEnvironmentAwsClient(envID))
		if err != nil {
			log.Errorf("error trying to start env %s: %v", envID, err)
		} else {
			log.Infof("successfully started env %s", envID)
		}
	} else {
		err = fmt.Errorf("env [%s] has no associated instances", envID)
		log.Errorf("env [%s] has no associated instances", envID)
	}
	return
}

// starts up an instance based on internal id (not aws instance id)
func toggleInstance(id, desiredState string) (response []byte, err error) {
	// use the mock function if enabled
	if MockEnabled {
		return mockToggleInstance(id, desiredState)
	}

	// validate desiredState
	if desiredState != "start" && desiredState != "stop" {
		err = fmt.Errorf("invalid desired state: %s", desiredState)
		return
	}
	// get the AWS instance id
	awsInstanceID := getAWSInstanceID(id)
	if awsInstanceID != "" {
		response, err = toggleInstances([]string{awsInstanceID}, desiredState, getInstanceAwsClient(id))
		if err != nil {
			log.Errorf("error trying to %s instance %s: %v", desiredState, id, err)
		} else {
			log.Infof("successfully toggled instance state (%s): %s", desiredState, id)
		}
	} else {
		err = fmt.Errorf("no mapping found between internal id (%s) and aws instance id", id)
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
func getEnvironmentAwsClient(envID string) *ec2.EC2 {
	for _, env := range cachedTable {
		if env.ID == envID {
			return awsClients[env.Region]
		}
	}
	return nil
}

// returns awsClient for the specific environment ID
func getInstanceAwsClient(instanceID string) *ec2.EC2 {
	for _, env := range cachedTable {
		for _, instance := range env.Instances {
			if instance.ID == instanceID {
				return awsClients[instance.Region]
			}
		}
	}
	return nil
}

// given an aws-power-toggle id, it will return the actual aws instance id
func getAWSInstanceID(id string) (awsInstanceID string) {
	for _, env := range cachedTable {
		for _, instance := range env.Instances {
			if instance.ID == id {
				awsInstanceID = instance.InstanceID
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
