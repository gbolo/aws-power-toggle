package backend

import (
	"testing"
)

// TestNewRouter checks that all defined routes are loaded properly
func TestNewRouter(t *testing.T) {
	// check that routes has entries
	if len(routes) == 0 {
		t.Error("routes are missing!")
	}

	// validate that we load all routes into the router
	router := newRouter()
	for _, route := range routes {
		if r := router.GetRoute(route.Name); r == nil {
			log.Errorf("route with name %s was not loaded correctly", route.Name)
		}
	}
}

// ensure that the correct route endpoints are being returned
func TestGetEndpoint(t *testing.T) {
	expectedResult := "/api/v1/testroute"
	if getEndpoint("testroute") != expectedResult {
		t.Errorf("api endpoint naming logic has changed: %s != %s", getEndpoint("testroute"), expectedResult)
	}
}
