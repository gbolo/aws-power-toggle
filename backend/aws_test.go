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
	envID := "356f6265efcc"

	if len(cachedTable) == 0 {
		t.Fatalf("cachedTable is of 0 size!")
	}

	env, found := getEnvironmentByID(envID)
	if !found {
		t.Errorf("unable to retrieve test env %s", envID)
	}
	if env.Name != "mockenv1" && len(env.Instances) < 1 {
		t.Errorf("unexpected env retrieved")
	}
	if _, found := getEnvironmentByID("incorrectEnvId"); found {
		t.Errorf("getEnvironmentByID reported found for an invalid env id")
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
	if cachedTable[1].State != EnvStateMixed {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, EnvStateMixed)
	}

	// set an instance in this env to pending, we should see env state change
	cachedTable[1].Instances[1].State = "pending"
	updateEnvDetails()
	if cachedTable[1].State != EnvStateChanging {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, EnvStateChanging)
	}

	// set all instances to running, we should see env state change
	for i := range cachedTable[1].Instances {
		cachedTable[1].Instances[i].State = "running"
	}
	updateEnvDetails()
	if cachedTable[1].State != EnvStateRunning {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, EnvStateRunning)
	}

	// set all instances to stopped, we should see env state change
	for i := range cachedTable[1].Instances {
		cachedTable[1].Instances[i].State = "stopped"
	}
	updateEnvDetails()
	if cachedTable[1].State != EnvStateStopped {
		t.Errorf("unexpected env state: got %s expected %s", cachedTable[1].State, EnvStateStopped)
	}

	// ensure that our counts are functioning correctly
	currentTotalCPUCount := cachedTable[1].TotalVCPU
	cachedTable[1].Instances[1].VCPU++
	updateEnvDetails()
	if (currentTotalCPUCount + 1) != cachedTable[1].TotalVCPU {
		t.Errorf("Total vCPU count has an unexpected value: %d", cachedTable[1].TotalVCPU)
	}
}

func TestToggleInstance(t *testing.T) {
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}

	// test toggleInstance
	for desiredState, actualState := range map[string]string{
		"stop":  "stopped",
		"start": "running",
	} {
		_, err := toggleInstance(cachedTable[1].Instances[1].ID, desiredState)
		if err != nil || cachedTable[1].Instances[1].State != actualState {
			t.Errorf("go %s but wanted %s. Err: %v", cachedTable[1].Instances[1].State, actualState, err)
		}
	}
}

func TestGetAWSInstanceId(t *testing.T) {
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}

	if getAWSInstanceID(cachedTable[1].Instances[1].ID) != cachedTable[1].Instances[1].InstanceID {
		t.Errorf("getAWSInstanceID is not returning correct id")
	}
}

func TestEnvStartStop(t *testing.T) {
	// disabled mock delays and chance of errors
	unitTestRunning = true

	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}
	envID := "4f9f1afb29f1"

	_, err := startupEnv(envID)
	if err != nil {
		t.Errorf("startupEnv return and error: %v", err)
	}
	updateEnvDetails()
	state, _ := getEnvState(envID)
	if state != "running" {
		t.Errorf("test env is not in running state: %s", state)
	}

	_, err = shutdownEnv(envID)
	if err != nil {
		t.Errorf("startupEnv return and error: %v", err)
	}
	updateEnvDetails()
	state, _ = getEnvState(envID)
	if state != "stopped" {
		t.Errorf("test env is not in stopped state: %s", state)
	}
}

func resetMockData() error {
	// disabled mock delays and chance of errors
	unitTestRunning = true
	mockEnabled = true
	cachedTable = cachedTable[:0]
	return mockRefreshTable()
}

func getEnvState(envID string) (string, bool) {
	env, found := getEnvironmentByID(envID)
	if !found {
		return "", found
	}
	return env.State, found
}
