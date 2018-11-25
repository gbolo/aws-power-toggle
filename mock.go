package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// mock of refreshTable
func mockRefreshTable() (err error) {
	// we only need to load initial test data when cachedTable is empty
	if len(cachedTable) == 0 {
		cachedTableJsonFile, err := os.Open("testdata/mock/mock_env_cachedTable.json")
		if err != nil {
			log.Fatalf("mock API is enabled, but can't load test data: %s", err)
		}
		defer cachedTableJsonFile.Close()
		cachedTableBytes, _ := ioutil.ReadAll(cachedTableJsonFile)
		err = json.Unmarshal(cachedTableBytes, &cachedTable)
		if err != nil {
			log.Fatalf("mock API is enabled, but can't unmarshal json file: %s", err)
		}
	}

	updateEnvDetails()
	log.Debugf("MOCK: valid environment(s) in cache: %d", len(cachedTable))
	return
}

// mock of shutdownEnv
func mockShutdownEnv(envId string) (response []byte, err error) {
	instanceIds := getInstanceIds(envId, "running")
	if len(instanceIds) > maxInstancesToShutdown {
		err = fmt.Errorf("SAFETY: env [%s] has too many associated instances to shutdown %d", envId, len(instanceIds))
		log.Debugf("SAFETY: instances: %v", instanceIds)
	} else if len(instanceIds) > 0 {
		// set all instances to stopped
		for e, env := range cachedTable {
			if env.Id == envId {
				for i, _ := range cachedTable[e].Instances {
					cachedTable[e].Instances[i].State = "stopped"
				}
				break
			}
		}
		response = []byte(`{"mock": "OK"}`)
		log.Infof("MOCK: successfully stopped env %s", envId)
	} else {
		err = fmt.Errorf("MOCK: env [%s] has no associated instances", envId)
		log.Errorf("env [%s] has no associated instances", envId)
	}
	return
}

// mock of startupEnv
func mockStartupEnv(envId string) (response []byte, err error) {
	instanceIds := getInstanceIds(envId, "stopped")
	if len(instanceIds) > 0 {
		// set all instances to running
		for e, env := range cachedTable {
			if env.Id == envId {
				for i, _ := range cachedTable[e].Instances {
					cachedTable[e].Instances[i].State = "running"
				}
				break
			}
		}
		response = []byte(`{"mock": "OK"}`)
		log.Infof("MOCK: successfully started env %s", envId)
	} else {
		err = fmt.Errorf("MOCK: env [%s] has no associated instances", envId)
		log.Errorf("MOCK: env [%s] has no associated instances", envId)
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
	// get the AWS instance id
	awsInstanceId := getAWSInstanceId(id)
	if awsInstanceId != "" {
		// set instance to desired state
		for e, env := range cachedTable {
			for i, instance := range env.Instances {
				if instance.Id == id {
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
