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
