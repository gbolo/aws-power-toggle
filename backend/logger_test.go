package backend

import (
	"strings"
	"testing"

	logging "github.com/op/go-logging"
)

func TestLoggingInit(t *testing.T) {

	// test log level
	for _, setLevel := range []string{"ERROR", "INFO", "warning", "WARNING"} {
		loggingInit(setLevel)
		appliedLevel := logging.GetLevel("log_backend_f").String()
		if appliedLevel != strings.ToUpper(setLevel) {
			t.Errorf("loging level was not applied properly: %s != %s", appliedLevel, setLevel)
		}
	}
	// test invalid log level. Should get set to ERROR
	loggingInit("INVALID")
	if logging.GetLevel("log_backend_f").String() != "ERROR" {
		t.Errorf(
			"expected an invalid log level to result in ERROR but got: %s",
			logging.GetLevel("log_backend_f").String(),
		)
	}
}
