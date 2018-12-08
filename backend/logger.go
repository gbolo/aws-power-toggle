package backend

import (
	"os"

	logging "github.com/op/go-logging"
)

// logFormat is an RFC 5424 style log format
const logFormat = `%{level:.1s} %{time:2006-01-02 15:04:05} %{program}[%{pid}]: %{id:05d} %{shortfile} %{shortfunc} %{message}`

// global logger for this package
var log = logging.MustGetLogger("aws-power-toggle")

func loggingInit(logLevelString string) {

	// In case of an invalid log level we default to ERROR
	var logLevel = logging.ERROR

	// Validating log levels
	_, err := logging.LogLevel(logLevelString)
	if err != nil {
		log.Errorf("The specified logLevel (%s) is invalid, defaulting to ERROR!", logLevelString)
	} else {
		// Converting the logLevel to actual Level (int)
		logLevel, _ = logging.LogLevel(logLevelString)
	}

	// Special logging format for the logs
	var logFormatter logging.Formatter

	// RFC 5424 style log format
	logFormatter = logging.MustStringFormatter(logFormat)

	// Configure log backends
	logBackend := logging.NewLogBackend(os.Stdout, "", 0)

	// Bind log formats and log backends together
	logBackendFormatter := logging.NewBackendFormatter(logBackend, logFormatter)

	// Setup the log level for each backend
	leveledBackend := logging.AddModuleLevel(logBackendFormatter)
	leveledBackend.SetLevel(logLevel, "")

	logging.SetBackend(leveledBackend)
}
