package backend

import (
	"fmt"
	"os"
	"testing"
)

func TestConfigInit(t *testing.T) {
	maxInstancesToShutdown = -1

	testConfig := "../testdata/sampleconfig"
	goPath := os.Getenv("GOPATH")
	if goPath != "" {
		testConfig = fmt.Sprintf("%s/src/github.com/gbolo/aws-power-toggle/testdata/sampleconfig", goPath)
	}
	ConfigInit(testConfig, true)
	if maxInstancesToShutdown == -1 {
		t.Error("ConfigInit did not change maxInstancesToShutdown")
	}
}
