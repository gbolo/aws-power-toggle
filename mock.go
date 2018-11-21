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
		cachedTableJsonFile, err := os.Open("testdata/mock_env_cachedTable.json")
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
func mockShutdownEnv(envName string) (response []byte, err error) {
	instanceIds := getInstanceIds(envName)
	if len(instanceIds) > maxInstancesToShutdown {
		err = fmt.Errorf("SAFETY: env [%s] has too many associated instances to shutdown %d", envName, len(instanceIds))
		log.Debugf("SAFETY: instances: %v", instanceIds)
	} else if len(instanceIds) > 0 {
		// set all instances to stopped
		for e, env := range cachedTable {
			if env.Name == envName {
				for i, _ := range cachedTable[e].Instances {
					cachedTable[e].Instances[i].State = "stopped"
				}
				break
			}
		}
		response = []byte(`{"mock": "OK"}`)
		log.Infof("MOCK: successfully stopped env %s", envName)
	} else {
		err = fmt.Errorf("MOCK: env [%s] has no associated instances", envName)
		log.Errorf("env [%s] has no associated instances", envName)
	}
	return
}

// mock of startupEnv
func mockStartupEnv(envName string) (response []byte, err error) {
	instanceIds := getInstanceIds(envName)
	if len(instanceIds) > 0 {
		// set all instances to running
		for e, env := range cachedTable {
			if env.Name == envName {
				for i, _ := range cachedTable[e].Instances {
					cachedTable[e].Instances[i].State = "running"
				}
				break
			}
		}
		response = []byte(`{"mock": "OK"}`)
		log.Infof("MOCK: successfully started env %s", envName)
	} else {
		err = fmt.Errorf("MOCK: env [%s] has no associated instances", envName)
		log.Errorf("MOCK: env [%s] has no associated instances", envName)
	}
	return
}
