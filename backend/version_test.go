package backend

import "testing"

func TestGetVersion(t *testing.T) {

	expectedResult := `{"version":"devel","git_hash":"unknown","build_date":"unknown"}`
	if getVersionResponse() != expectedResult {
		t.Error("version output has changed")
	}
}
