package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// mock of refreshTable
func mockRefreshTable() (err error) {
	// we only need to load initial test data when cachedTable is empty
	if len(cachedTable) == 0 {
		cachedTableJSONFile, err := os.Open("../testdata/mock/mock_env_cachedTable.json")
		if err != nil {
			log.Fatalf("mock API is enabled, but can't load test data: %s", err)
		}
		defer cachedTableJSONFile.Close()
		cachedTableBytes, _ := ioutil.ReadAll(cachedTableJSONFile)
		err = json.Unmarshal(cachedTableBytes, &cachedTable)
		if err != nil {
			log.Fatalf("mock API is enabled, but can't unmarshal json file: %s", err)
		}
	}
	if ExperimentalEnabled {
		calculateEnvBills()
	}
	cachedTableTemp := cachedTable
	cachedTable = cachedTable[:0]
	for _, env := range cachedTableTemp {
		for _, instanceObj := range env.Instances {
			// determine instance cpu and memory
			if details, found := getInstanceTypeDetails(instanceObj.InstanceType); found {
				instanceObj.MemoryGB = details.MemoryGB
				instanceObj.VCPU = details.VCPU
				if pricingstr, ok := details.PricingHourlyByRegion["ca-central-1"]; ok {
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
	log.Debugf("MOCK: valid environment(s) in cache: %d", len(cachedTable))
	return
}

// mock of shutdownEnv
func mockShutdownEnv(envID string) (response []byte, err error) {
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

// mock toggleInstance
func mockToggleInstance(id, desiredState string) (response []byte, err error) {
	// validate desiredState
	if desiredState != "start" && desiredState != "stop" {
		err = fmt.Errorf("invalid desired state: %s", desiredState)
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
