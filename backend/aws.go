package backend

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/liip/sheriff"
	"github.com/spf13/viper"
)

const (
	// defines environment states
	// means that ALL instances for an env are in "running" state
	ENV_RUNNING = "running"
	// means that ALL instances for an env are in "stopped" state
	ENV_DOWN = "stopped"
	// means that instances for an env are in EITHER "stopped" or "running" state
	ENV_MIXED = "mixed"
	// means that AT LEAST ONE instance for an env is NOT in a "running" state or "stopped" state
	ENV_CHANGING = "changing"
)

var (
	// global aws client
	awsClient *ec2.EC2
	// global cached env list
	cachedTable envList
	// lock to prevent concurrent refreshes
	cachedTableLock sync.Mutex
	// aws region
	awsRegion string
	// aws tags
	requiredTagKey, requiredTagValue, environmentTagKey string
	// safety, will refuse to shutdown if more than this amount of instances is requested
	maxInstancesToShutdown int
	// ignore these instance types
	instanceTypeIgnore []string
	// ignore these environment names
	envNameIgnore []string

	// enables mocking of API calls to aws for development purposes
	MOCK_ENABLED = false
)

type virtualMachine struct {
	// Id unique to this application
	Id string `json:"id" groups:"summary,details"`

	// these values are straight from aws api
	InstanceId   string `json:"instance_id" groups:"summary,details"`
	InstanceType string `json:"instance_type" groups:"summary,details"`
	Name         string `json:"name" groups:"summary,details"`
	State        string `json:"state" groups:"summary,details"`
	Environment  string `json:"environment" groups:"summary,details"`

	// these values are mapped from another source for aws
	VCPU     int     `json:"vcpu" groups:"summary,details"`
	MemoryGB float32 `json:"memory_gb" groups:"summary,details"`
}

type environment struct {
	// Id unique to this application
	Id        string           `json:"id" groups:"summary,details"`
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
		cachedTable[i].Region = awsRegion

		// compute a unique identifier for this environment
		cachedTable[i].Id = ComputeId(
			cachedTable[i].Provider,
			cachedTable[i].Region,
			env.Name,
		)

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
			//   InstanceId is already unique, but this will make ids consistent
			//   in case we add other cloud providers
			cachedTable[i].Instances[c].Id = ComputeId(
				cachedTable[i].Provider,
				cachedTable[i].Region,
				instance.InstanceId,
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
			cachedTable[i].State = ENV_RUNNING
		case cachedTable[i].TotalInstances == cachedTable[i].StoppedInstances:
			cachedTable[i].State = ENV_DOWN
		case cachedTable[i].TotalInstances == (cachedTable[i].RunningInstances + cachedTable[i].StoppedInstances):
			cachedTable[i].State = ENV_MIXED
		default:
			cachedTable[i].State = ENV_CHANGING
		}
	}
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
		}
		cachedTable = append(cachedTable, ec2env)
	}
}

// polls aws for updates to cachedTable
func refreshTable() (err error) {
	cachedTableLock.Lock()
	defer cachedTableLock.Unlock()

	// use the mock function if enabled
	if MOCK_ENABLED {
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

	req := awsClient.DescribeInstancesRequest(params)
	resp, err := req.Send()
	if err != nil {
		log.Errorf("failed to describe instances, %s, %v", awsRegion, err)
		return
	}
	log.Infof("aws poll was successful, clearing old cached table")
	cachedTable = cachedTable[:0]

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			instanceObj := virtualMachine{InstanceId: *instance.InstanceId, State: string(instance.State.Name), InstanceType: string(instance.InstanceType)}
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
			}
			if validateEnvName(instanceObj.Environment) {
				addInstance(&instanceObj)
			}
		}
	}
	updateEnvDetails()
	log.Debugf("valid environment(s) in cache: %d", len(cachedTable))
	return
}

// get instance ids for an environment with a specific state
// this is used for power up/down commands against aws API
func getInstanceIds(envId, state string) (instanceIds []string) {
	for _, env := range cachedTable {
		if env.Id == envId {
			for _, instance := range env.Instances {
				if instance.State == state {
					instanceIds = append(instanceIds, instance.InstanceId)
				}
			}
		}
	}
	return
}

// toggleInstances can start or stop a list of instances
func toggleInstances(instanceIds []string, desiredState string) (response []byte, err error) {
	if len(instanceIds) < 1 {
		err = fmt.Errorf("no instanceIds have been provided")
		return
	}

	// supported states are: start, stop
	switch desiredState {
	case "start":
		input := &ec2.StartInstancesInput{
			InstanceIds: instanceIds,
			DryRun:      aws.Bool(false),
		}

		req := awsClient.StartInstancesRequest(input)
		awsResponse, reqErr := req.Send()
		response, _ = json.MarshalIndent(awsResponse, "", "  ")
		err = reqErr
		return

	case "stop":
		input := &ec2.StopInstancesInput{
			InstanceIds: instanceIds,
			DryRun:      aws.Bool(false),
		}

		req := awsClient.StopInstancesRequest(input)
		awsResponse, reqErr := req.Send()
		response, _ = json.MarshalIndent(awsResponse, "", "  ")
		err = reqErr
		return

	default:
		err = fmt.Errorf("unsupported desiredState soecified")
		return
	}
}

// shuts down an env
func shutdownEnv(envId string) (response []byte, err error) {
	// use the mock function if enabled
	if MOCK_ENABLED {
		return mockShutdownEnv(envId)
	}

	instanceIds := getInstanceIds(envId, "running")
	if len(instanceIds) > maxInstancesToShutdown {
		err = fmt.Errorf("SAFETY: env [%s] has too many associated instances to shutdown %d", envId, len(instanceIds))
		log.Debugf("SAFETY: instances: %v", instanceIds)
	} else if len(instanceIds) > 0 {
		response, err = toggleInstances(instanceIds, "stop")
		if err != nil {
			log.Errorf("error trying to stop env %s: %v", envId, err)
		} else {
			log.Infof("successfully stopped env %s", envId)
		}
	} else {
		err = fmt.Errorf("env [%s] has no associated instances", envId)
		log.Errorf("env [%s] has no associated instances", envId)
	}
	return
}

// starts up an env
func startupEnv(envId string) (response []byte, err error) {
	// use the mock function if enabled
	if MOCK_ENABLED {
		return mockStartupEnv(envId)
	}

	instanceIds := getInstanceIds(envId, "stopped")
	if len(instanceIds) > 0 {
		response, err = toggleInstances(instanceIds, "start")
		if err != nil {
			log.Errorf("error trying to start env %s: %v", envId, err)
		} else {
			log.Infof("successfully started env %s", envId)
		}
	} else {
		err = fmt.Errorf("env [%s] has no associated instances", envId)
		log.Errorf("env [%s] has no associated instances", envId)
	}
	return
}

// starts up an instance based on internal id (not aws instance id)
func toggleInstance(id, desiredState string) (response []byte, err error) {
	// use the mock function if enabled
	if MOCK_ENABLED {
		return mockToggleInstance(id, desiredState)
	}

	// validate desiredState
	if desiredState != "start" && desiredState != "stop" {
		err = fmt.Errorf("invalid desired state: %s", desiredState)
		return
	}
	// get the AWS instance id
	awsInstanceId := getAWSInstanceId(id)
	if awsInstanceId != "" {
		response, err = toggleInstances([]string{awsInstanceId}, desiredState)
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
func getEnvironmentById(envId string) (environment, bool) {
	for _, env := range cachedTable {
		if env.Id == envId {
			return env, true
		}
	}
	return environment{}, false
}

// given an aws-power-toggle id, it will return the actual aws instance id
func getAWSInstanceId(id string) (awsInstanceId string) {
	for _, env := range cachedTable {
		for _, instance := range env.Instances {
			if instance.Id == id {
				awsInstanceId = instance.InstanceId
				break
			}
		}
	}
	return
}

// getMarshalledRespone will filter out fields from the struct based on predefined groups
func getMarshalledRespone(data interface{}, groups ...string) (response []byte, err error) {
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

	pollInterval := time.Minute * time.Duration(viper.GetInt("aws.polling_interval"))
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
