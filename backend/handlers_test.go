// handlers_test.go
package backend

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpHandlers(t *testing.T) {

	// disabled mock delays and chance of errors
	unitTestRunning = true

	// discard logs
	loggingInit("INFO")

	// prepare data
	if err := resetMockData(); err != nil {
		t.Fatalf("mockRefreshTable failed: %v", err)
	}

	type req struct {
		method   string
		endpoint string
		status   int
	}

	// ONLY check status code for now...
	// TODO: add checks for response body
	for _, testReq := range []req{
		{"GET", "/", http.StatusNotFound},
		{"GET", getEndpoint("version"), http.StatusOK},
		{"GET", getEndpoint("config"), http.StatusOK},
		{"POST", getEndpoint("refresh"), http.StatusOK},
		{"GET", getEndpoint("env/summary"), http.StatusOK},
		{"GET", getEndpoint("env/details"), http.StatusOK},
		{"GET", getEndpoint("env/4f9f1afb29f1/summary"), http.StatusOK},
		{"GET", getEndpoint("env/4f9f1afb29f1/details"), http.StatusOK},
		{"GET", getEndpoint("env/invalid/summary"), http.StatusNotFound},
		{"GET", getEndpoint("env/invalid/details"), http.StatusNotFound},
		{"POST", getEndpoint("env/4f9f1afb29f1/start"), http.StatusOK},
		{"POST", getEndpoint("env/4f9f1afb29f1/stop"), http.StatusOK},
		{"POST", getEndpoint("env/invalid/start"), http.StatusInternalServerError},
		{"POST", getEndpoint("env/invalid/stop"), http.StatusInternalServerError},
		{"POST", getEndpoint("instance/906d663b6ecd/start"), http.StatusOK},
		{"POST", getEndpoint("instance/906d663b6ecd/stop"), http.StatusOK},
		{"POST", getEndpoint("instance/invalid/start"), http.StatusInternalServerError},
		{"POST", getEndpoint("instance/invalid/stop"), http.StatusInternalServerError},
	} {

		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(testReq.method, testReq.endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		// since we use gorilla mux for our handers, we need to pass the request over to that
		mux := newRouter()
		mux.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != testReq.status {
			t.Errorf("handler returned wrong status code for endpoint: %s: got %v want %v",
				testReq.endpoint, status, testReq.status)
		}
	}
}
