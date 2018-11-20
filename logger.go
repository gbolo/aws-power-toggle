package main

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("aws-power-toggle")

func loggingInit(loglevel string) {

	// In case of an invalid log level we default to ERROR
	var log_level logging.Level = logging.ERROR

	// Validating log levels
	_, err := logging.LogLevel(loglevel)
	if err != nil {
		log.Errorf("The specified log_level (%s) is invalid, defaulting to ERROR!", loglevel)
	} else {
		// Converting the log_level to actual Level (int)
		log_level, _ = logging.LogLevel(loglevel)
	}

	// Special logging format for the logs
	var log_format logging.Formatter

	// RFC 5424 style log format
	log_format = logging.MustStringFormatter(
		`%{level:.1s} ` +
			`%{time:2006-01-02 15:04:05} ` +
			`%{program}[%{pid}]: ` +
			`%{id:05d} ` +
			`%{shortfile} ` +
			`%{shortfunc} ` +
			`%{message}`)

	// Configure log backends
	log_backend := logging.NewLogBackend(os.Stdout, "", 0)

	// Bind log formats and log backends together
	log_backend_f := logging.NewBackendFormatter(log_backend, log_format)

	// Setup the log level for each backends
	log_backend_fl := logging.AddModuleLevel(log_backend_f)
	log_backend_fl.SetLevel(log_level, "")

	logging.SetBackend(log_backend_fl)
}
