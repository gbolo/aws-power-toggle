package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// returns all envs
func handlerEnv(w http.ResponseWriter, req *http.Request) {

	jData, err := json.Marshal(cachedTable)
	if err != nil {
		log.Errorf("error parsing json: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

// returns summary of all envs
func handlerEnvSummary(w http.ResponseWriter, req *http.Request) {

	// refresh the data
	if err := refreshTable(); err != nil {
		log.Errorf("refresh error: %v", err)
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
		return
	}

	// return result
	jData, err := json.Marshal(getEnvSummary())
	if err != nil {
		log.Errorf("error parsing json: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

// handler to powerdown an env
func handlerEnvPowerdown(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// get vars from request to determine environment
	vars := mux.Vars(req)
	envName := vars["env"]

	if envName != "" {
		res, err := shutdownEnv(envName)
		writeJsonResponse(w, err, res)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"empty environment name\"}\n")
	}
}

// handler to start up an env
func handlerEnvStartup(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// get vars from request to determine environment
	vars := mux.Vars(req)
	envName := vars["env"]
	log.Infof("starting env: %s", envName)

	if envName != "" {
		res, err := startupEnv(envName)
		writeJsonResponse(w, err, res)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"empty environment name\"}\n")
	}
}

// handler to refresh envs
func handlerEnvRefresh(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := refreshTable(); err != nil {
		log.Errorf("refresh error: %v", err)
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
	} else {
		log.Info("refresh successful")
		fmt.Fprint(w, "{\"status\":\"OK\"}\n")
	}
}

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
