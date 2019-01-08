package backend

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

// unitTestRunning lets use know if a unit test is running
// currently, it will disable introduced delays and chance of mocked errors
var unitTestRunning = false

// mock of shutdownEnv
func mockShutdownEnv(envID string) (response []byte, err error) {
	// introduce delays and possible error
	err = mockDelayWithPossibleError()
	if err != nil {
		response = []byte(fmt.Sprintf(`{"error":"%v"}`, err))
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
		// MOCK BILLING: update toggled off instances map
		if ExperimentalEnabled {
			putToggledOffInstanceIDs(instanceIds)
		}
		response = []byte(`{"mock": "OK"}`)
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
	err = mockDelayWithPossibleError()
	if err != nil {
		response = []byte(fmt.Sprintf(`{"error":"%v"}`, err))
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
		// MOCK BILLING: update toggled off instances map
		if ExperimentalEnabled {
			deleteToggledOffInstanceIDs(instanceIds)
		}
		response = []byte(`{"mock": "OK"}`)
		log.Infof("MOCK: successfully started env %s", envID)
	} else {
		err = fmt.Errorf("MOCK: env [%s] has no associated instances", envID)
		log.Errorf("MOCK: env [%s] has no associated instances", envID)
	}
	return
}

// mock of toggleInstance
func mockToggleInstance(id, desiredState string) (response []byte, err error) {
	// validate desiredState
	if desiredState != "start" && desiredState != "stop" {
		err = fmt.Errorf("invalid desired state: %s", desiredState)
		return
	}
	// introduce delays and possible error
	err = mockDelayWithPossibleError()
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
						// MOCK BILLING: update toggled off instances map
						if ExperimentalEnabled {
							deleteToggledOffInstanceIDs([]string{instance.InstanceID})
						}
					case "stop":
						cachedTable[e].Instances[i].State = "stopped"
						// MOCK BILLING: update toggled off instances map
						if ExperimentalEnabled {
							putToggledOffInstanceIDs([]string{instance.InstanceID})
						}
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
func mockDelayWithPossibleError() (err error) {
	// if we are doing unit tests, this should be disabled
	if unitTestRunning {
		return
	}

	// random delay betwen 100-2100 ms
	r := rand.Intn(2000) + 100
	if viper.GetBool("mock.delay") {
		time.Sleep(time.Duration(r) * time.Millisecond)
	}

	// 1/4 chance of producing an error
	if r%4 == 0 && viper.GetBool("mock.errors") {
		err = fmt.Errorf("MOCK: Fate has thrown you an error")
	}
	return
}
