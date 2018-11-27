package backend

import "testing"

func TestLoadAwsInstanceDetailsJson(t *testing.T) {

	// this should NEVER return an error
	// if it does then the json data may be incorrect
	if err := loadAwsInstanceDetailsJson(); err != nil {
		t.Error("loadAwsInstanceDetailsJson returned an error")
	}

	// check if instanceTypeDetailsCache was updated properly
	if len(instanceTypeDetailsCache) < 1 {
		t.Error("instanceTypeDetailsCache has not been updated")
	}
}