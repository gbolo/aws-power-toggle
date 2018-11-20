package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// all defined api routes
var routes = Routes{

	// API endpoints
	Route{
		"Environments",
		"GET",
		"/api/env",
		handlerEnv,
	},
	Route{
		"Environments",
		"GET",
		"/api/env/summary",
		handlerEnvSummary,
	},
	Route{
		"Refresh",
		"POST",
		"/api/env/refresh",
		handlerEnvRefresh,
	},
	Route{
		"StartUp",
		"POST",
		"/api/env/startup/{env}",
		handlerEnvStartup,
	},
	Route{
		"PowerDown",
		"POST",
		"/api/env/powerdown/{env}",
		handlerEnvPowerdown,
	},

	// UI endpoints
	Route{
		"Main",
		"GET",
		"/",
		handlerMain,
	},
}

func newRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc

		// add routes to mux
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// add route to mux to handle static files
	staticPath := viper.GetString("server.static_files_dir")
	if staticPath == "" {
		staticPath = "./static"
	}

	router.
		Methods("GET").
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))

	return router
}
