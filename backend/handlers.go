package backend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// returns version information
func handlerVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, getVersionResponse())
}

// handler for all environments
func handlerEnvAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// refresh the data
	if err := refreshTable(); err != nil {
		log.Errorf("refresh error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
		return
	}

	// get vars from request to determine if environment id was specified
	vars := mux.Vars(req)
	group := vars["group"]

	envAllResponse := struct {
		EnvList           envList `json:"envList" groups:"summary,details"`
		TotalBillsAccrued string  `json:"totalBillsAccrued" groups:"summary,details"`
		TotalBillsSaved   string  `json:"totalBillsSaved" groups:"summary,details"`
	}{
		EnvList:           cachedTable,
		TotalBillsAccrued: fmt.Sprintf("%.02f", totalBillsAccrued),
		TotalBillsSaved:   fmt.Sprintf("%.02f", totalBillsSaved),
	}

	// prepare result and return it
	if response, err := getMarshaledResponse(envAllResponse, group); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
	} else {
		w.Write(response)
	}
}

// handler for single environment
func handlerEnvSingle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// refresh the data
	if err := refreshTable(); err != nil {
		log.Errorf("refresh error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
		return
	}

	// get vars from request to determine if environment id was specified
	vars := mux.Vars(req)
	envID := vars["env-id"]
	group := vars["group"]

	// filter this environment id
	envData, found := getEnvironmentByID(envID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\":\"environment not found\"}\n")
		return
	}
	response, err := getMarshaledResponse(envData, group)

	// return filtered result
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
	} else {
		w.Write(response)
	}
}

// handler for power toggling an environment
func handlerEnvPowerToggle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get vars from request to determine environment
	vars := mux.Vars(req)
	envID := vars["env-id"]
	state := vars["state"]

	switch state {
	case "start":
		response, err := startupEnv(envID)
		writeJSONResponse(w, err, response)
	case "stop":
		response, err := shutdownEnv(envID)
		writeJSONResponse(w, err, response)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"invalid request\"}\n")
	}
}

// handler for power toggling an instance
func handlerInstancePowerToggle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get vars from request to determine environment
	vars := mux.Vars(req)
	id := vars["instance-id"]
	state := vars["state"]

	if state == "start" || state == "stop" {
		response, err := toggleInstance(id, state)
		writeJSONResponse(w, err, response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"invalid request\"}\n")
	}
}

// handler to refresh cache
func handlerRefresh(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := refreshTable(); err != nil {
		log.Errorf("refresh error: %v", err)
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
	} else {
		log.Info("refresh successful")
		fmt.Fprint(w, "{\"status\":\"OK\"}\n")
	}
}

// wrapper for json responses with error support
func writeJSONResponse(w http.ResponseWriter, err error, response []byte) {
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		if len(response) > 0 {
			w.Write(response)
		} else {
			fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
		}
	}
}
