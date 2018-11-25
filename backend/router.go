package backend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

const (
	ApiVersion     = "1"
	endpointFormat = "/api/v%s/%s"
)

// getEndpoint returns a properly formatted API endpoint
func getEndpoint(suffix string) string {
	return fmt.Sprintf(endpointFormat, ApiVersion, suffix)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// all defined server endpoints
var routes = Routes{

	// API endpoints
	Route{
		"Version",
		"GET",
		getEndpoint("version"),
		handlerVersion,
	},
	Route{
		"Refresh",
		"POST",
		getEndpoint("refresh"),
		handlerRefresh,
	},
	Route{
		"EnvAll",
		"GET",
		getEndpoint("env/{group:summary|details}"),
		handlerEnvAll,
	},
	Route{
		"EnvSingle",
		"GET",
		getEndpoint("env/{env-id}/{group:summary|details}"),
		handlerEnvSingle,
	},
	Route{
		"EnvPowerToggle",
		"POST",
		getEndpoint("env/{env-id}/{state:start|stop}"),
		handlerEnvPowerToggle,
	},
	Route{
		"InstancePowerToggle",
		"POST",
		getEndpoint("instance/{instance-id}/{state:start|stop}"),
		handlerInstancePowerToggle,
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
