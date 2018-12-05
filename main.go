package main

import (
	"flag"
	"os"

	"github.com/gbolo/aws-power-toggle/backend"
)

func main() {
	// parse flags
	cfgFile := flag.String("config", "", "path to config file")
	outputVersion := flag.Bool("version", false, "prints version then exits")
	flag.Parse()

	// print version and exit if flag is present
	if *outputVersion {
		backend.PrintVersion()
		os.Exit(0)
	}

	// start the Backend
	backend.StartBackendDeamon(*cfgFile)
}
