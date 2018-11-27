package backend

import (
	"os"
	"testing"
)

func TestCheckInstanceType(t *testing.T) {
	// init the test config
	os.Setenv("POWER_TOGGLE_AWS_IGNORE_INSTANCE_TYPES", "c5d.18xlarge c5d.9xlarge")
	ConfigInit("../testdata/sampleconfig/power-toggle-config.yaml", false)

	testCases := map[string]bool{
		"c5d.18xlarge": false,
		"c5d.9xlarge":  false,
		"c5.18xlarge":  true,
		"t2.medium":    true,
	}

	for input, expectedResult := range testCases {
		if checkInstanceType(input) != expectedResult {
			t.Errorf("input(%s) != expectedResult(%v)", input, expectedResult)
		}
	}
}

func TestValidateEnvName(t *testing.T) {
	// init the test config
	os.Setenv("POWER_TOGGLE_AWS_IGNORE_ENVIRONMENTS", "ignoredEnv1 ignoredEnv2")
	ConfigInit("../testdata/sampleconfig/power-toggle-config.yaml", false)

	testCases := map[string]bool{
		"ignoredEnv1": false,
		"ignoredEnv2": false,
		"env3":        true,
		"env4":        true,
	}

	for input, expectedResult := range testCases {
		if validateEnvName(input) != expectedResult {
			t.Errorf("input(%s) != expectedResult(%v)", input, expectedResult)
		}
	}
}

func TestGetEnvironmentById(t *testing.T) {
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}
	envId := "356f6265efcc"

	if len(cachedTable) == 0 {
		t.Fatalf("cachedTable is of 0 size!")
	}

	env, found := getEnvironmentById(envId)
	if !found {
		t.Errorf("unable to retrieve test env %s", envId)
	}
	if env.Name != "mockenv1" && len(env.Instances) < 1 {
		t.Errorf("unexpected env retrieved")
	}
	if _, found := getEnvironmentById("incorrectEnvId"); found {
		t.Errorf("getEnvironmentById reported found for an invalid env id")
	}
}


func TestUpdateEnvDetails(t *testing.T) {
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}

	// ensure our mocked json initial state is correct
	startingInstanceState := cachedTable[1].Instances[1].State
	if startingInstanceState != "stopped" {
		t.Error("our moc json has unexpected init values. Has it changed?")
	}
	startingEnvState := cachedTable[1].State
	if startingEnvState != "stopped" {
		t.Error("our moc json has unexpected init values. Has it changed?")
	}

	// set an instance in this env to running, we should see env state change
	cachedTable[1].Instances[1].State = "running"
	updateEnvDetails()
	if cachedTable[1].State != ENV_MIXED {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, ENV_MIXED)
	}

	// set an instance in this env to pending, we should see env state change
	cachedTable[1].Instances[1].State = "pending"
	updateEnvDetails()
	if cachedTable[1].State != ENV_CHANGING {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, ENV_CHANGING)
	}

	// set all instances to running, we should see env state change
	for i, _ := range cachedTable[1].Instances {
		cachedTable[1].Instances[i].State = "running"
	}
	updateEnvDetails()
	if cachedTable[1].State != ENV_RUNNING {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, ENV_RUNNING)
	}

	// set all instances to stopped, we should see env state change
	for i, _ := range cachedTable[1].Instances {
		cachedTable[1].Instances[i].State = "stopped"
	}
	updateEnvDetails()
	if cachedTable[1].State != ENV_DOWN {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, ENV_DOWN)
	}

	// ensure that our counts are functioning correctly
	currentTotalCpuCount := cachedTable[1].TotalVCPU
	cachedTable[1].Instances[1].VCPU++
	updateEnvDetails()
	if (currentTotalCpuCount + 1) != cachedTable[1].TotalVCPU {
		t.Errorf("Total vCPU count has an unexpected value: %d", cachedTable[1].TotalVCPU)
	}
}

func TestToggleInstance(t *testing.T) {
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}

	// test toggleInstance
	for desiredState, actualState := range map[string]string{
		"stop": "stopped",
		"start": "running",
	} {
		_, err := toggleInstance(cachedTable[1].Instances[1].Id, desiredState)
		if err != nil || cachedTable[1].Instances[1].State != actualState {
			t.Errorf("go %s but wanted %s. Err: %v", cachedTable[1].Instances[1].State, actualState, err)
		}
	}
}

func TestGetAWSInstanceId(t *testing.T) {
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}

	if getAWSInstanceId(cachedTable[1].Instances[1].Id) != cachedTable[1].Instances[1].InstanceId {
		t.Errorf("getAWSInstanceId is not returning correct id")
	}
}

func TestEnvStartStop(t *testing.T) {
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}
	envId := "4f9f1afb29f1"

	_, err := startupEnv(envId)
	if err != nil {
		t.Errorf("startupEnv return and error: %v", err)
	}
	updateEnvDetails()
	state, _ := getEnvState(envId)
	if state != "running" {
		t.Errorf("test env is not in running state: %s", state)
	}

	_, err = shutdownEnv(envId)
	if err != nil {
		t.Errorf("startupEnv return and error: %v", err)
	}
	updateEnvDetails()
	state, _ = getEnvState(envId)
	if state != "stopped" {
		t.Errorf("test env is not in stopped state: %s", state)
	}
}

func resetMockData() error {
	MOCK_ENABLED = true
	cachedTable = cachedTable[:0]
	return mockRefreshTable()
}

func getEnvState(envId string) (string, bool) {
	env, found := getEnvironmentById(envId)
	if !found {
		return "", found
	} else {
		return env.State, found
	}
}



