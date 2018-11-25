package backend

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

	// prepare result and return it
	if response, err := getMarshalledRespone(cachedTable, group); err != nil {
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
	envId := vars["env-id"]
	group := vars["group"]

	// filter this environment id
	envData, found := getEnvironmentById(envId)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\":\"environment not found\"}\n")
		return
	}
	response, err := getMarshalledRespone(envData, group)

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
	envId := vars["env-id"]
	state := vars["state"]

	switch state {
	case "start":
		response, err := startupEnv(envId)
		writeJsonResponse(w, err, response)
	case "stop":
		response, err := shutdownEnv(envId)
		writeJsonResponse(w, err, response)
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
		writeJsonResponse(w, err, response)
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
func writeJsonResponse(w http.ResponseWriter, err error, response []byte) {
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

// handler for main UI page
func handlerMain(w http.ResponseWriter, req *http.Request) {
	// write the templated HTML
	w.WriteHeader(http.StatusOK)
	writeTemplatedHtml(w, "main")

}

// write dashboard HTML from templates
func writeTemplatedHtml(w http.ResponseWriter, templateName string) {

	// read in all templates
	templateFiles, err := findAllTemplates()

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error reading template(s). See logs")
		log.Errorf("error geting template files: %s", err)
		return
	}

	// parse and add function to templates
	templates, err := template.New(templateName).
		Funcs(template.FuncMap{"StringsJoin": strings.Join}).
		ParseFiles(templateFiles...)

	if err == nil {
		templateData := map[string]string{
			"Region": viper.GetString("aws.region"),
		}
		execErr := templates.ExecuteTemplate(w, templateName, templateData)
		if execErr != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Error executing template(s). See logs")
			log.Errorf("template execute error: %s", execErr)
			return
		}

	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error parsing template(s). See logs")
		log.Errorf("template parse error: %s", err)
		return
	}

	return
}

// return list of all template filenames
func findAllTemplates() (templateFiles []string, err error) {

	files, err := ioutil.ReadDir("./www/templates")
	if err != nil {
		return
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".gohtml") {
			templateFiles = append(templateFiles, "./www/templates/"+filename)
		}
	}

	return
}
