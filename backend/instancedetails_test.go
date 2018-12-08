package backend

import (
	"testing"
)

func TestAwsInstanceDetails(t *testing.T) {

	// this should NEVER return an error
	// if it does then the json data may be incorrect
	if err := loadAwsInstanceDetailsJSON(); err != nil {
		t.Error("loadAwsInstanceDetailsJSON returned an error")
	}

	// check if instanceTypeDetailsCache was updated properly
	if len(instanceTypeDetailsCache) < 1 {
		t.Error("instanceTypeDetailsCache has not been updated")
	}

	// test that we can get instance type details
	for iType, expected := range map[string]bool{
		"t2.medium":   true,
		"c5d.4xlarge": true,
		"p99.invalid": false,
	} {
		_, got := getInstanceTypeDetails(iType)
		if expected != got {
			t.Errorf("for type: %s got %v, expected %v", iType, got, expected)
		}
	}
}
