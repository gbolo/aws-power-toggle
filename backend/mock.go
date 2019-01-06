package backend

import (
	"fmt"
	"math/rand"
	"time"
)

// unitTestRunning lets use know if a unit test is running
// currently, it will disable introduced delays and chance of mocked errors
var unitTestRunning = false

// mock of shutdownEnv
func mockShutdownEnv(envID string) (response []byte, err error) {
	// introduce delays and possible error
	err = mockDelayWithPossibleError(1)
	if err != nil {
		response = []byte(fmt.Sprintf(`{"error":"%s"}`, err))
		log.Errorf("mock error envID: %s: %s", envID, err)
		return
	}
	instanceIds := getInstanceIDs(envID, "running")
	if len(instanceIds) > maxInstancesToShutdown {
		err = fmt.Errorf("SAFETY: env [%s] has too many associated instances to shutdown %d", envID, len(instanceIds))
		log.Debugf("SAFETY: instances: %v", instanceIds)
	} else if len(instanceIds) > 0 {
		// set all instances to stopped
		for e, env := range cachedTable {
			if env.ID == envID {
				for i := range cachedTable[e].Instances {
					cachedTable[e].Instances[i].State = "stopped"
				}
				break
			}
		}
		response = []byte(fmt.Sprintf(`{"error":"%s"}`, err))
		log.Infof("MOCK: successfully stopped env %s", envID)
	} else {
		err = fmt.Errorf("MOCK: env [%s] has no associated instances", envID)
		log.Errorf("env [%s] has no associated instances", envID)
	}
	return
}

// mock of startupEnv
func mockStartupEnv(envID string) (response []byte, err error) {
	// introduce delays and possible error
	err = mockDelayWithPossibleError(1)
	if err != nil {
		response = []byte(fmt.Sprintf(`{"error":"%s"}`, err))
		log.Errorf("mock error envID: %s: %s", envID, err)
		return
	}
	instanceIds := getInstanceIDs(envID, "stopped")
	if len(instanceIds) > 0 {
		// set all instances to running
		for e, env := range cachedTable {
			if env.ID == envID {
				for i := range cachedTable[e].Instances {
					cachedTable[e].Instances[i].State = "running"
				}
				break
			}
		}
		response = []byte(`{"mock": "OK"}`)
		log.Infof("MOCK: successfully started env %s", envID)
	} else {
		err = fmt.Errorf("MOCK: env [%s] has no associated instances", envID)
		log.Errorf("MOCK: env [%s] has no associated instances", envID)
	}
	return
}

// mock
func mockToggleInstance(id, desiredState string) (response []byte, err error) {
	// validate desiredState
	if desiredState != "start" && desiredState != "stop" {
		err = fmt.Errorf("invalid desired state: %s", desiredState)
		return
	}
	// introduce delays and possible error
	err = mockDelayWithPossibleError(1)
	if err != nil {
		response = []byte(fmt.Sprintf(`{"error":"%s"}`, err))
		log.Errorf("mock error instance id: %s: %s", id, err)
		return
	}
	// get the AWS instance id
	awsInstanceID := getAWSInstanceID(id)
	if awsInstanceID != "" {
		// set instance to desired state
		for e, env := range cachedTable {
			for i, instance := range env.Instances {
				if instance.ID == id {
					switch desiredState {
					case "start":
						cachedTable[e].Instances[i].State = "running"
					case "stop":
						cachedTable[e].Instances[i].State = "stopped"
					}
					break
				}
			}
		}
		response = []byte(`{"mock": "OK"}`)
	} else {
		err = fmt.Errorf("no mapping found between internal id (%s) and aws instance id", id)
	}
	return
}

// mockDelayWithPossibleError will add a delay and possibly return an error.
// This is done to simulate real world delays and issues to aid in web UI development
func mockDelayWithPossibleError(delay int) (err error) {
	// if we are doing unit tests, this should be disabled
	if unitTestRunning {
		return
	}

	time.Sleep(time.Duration(delay) * time.Second)
	if rand.Intn(4) == 0 {
		err = fmt.Errorf("MOCK: Fate has thrown you an error")
	}
	return
}
