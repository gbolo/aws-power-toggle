package main

import (
	"flag"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// load instancetype details
	if err := loadAwsInstanceDetailsJson(); err != nil {
		log.Fatalf("could not load instance type details: %v", err)
	}

	// parse flags
	cfgFile := flag.String("config", "", "path to config file")
	outputVersion := flag.Bool("version", false, "prints version then exits")
	flag.Parse()

	// print version and exit if flag is present
	if *outputVersion {
		printVersion()
		os.Exit(0)
	}

	// init config and logging
	ConfigInit(*cfgFile)

	// start http server
	go StartServer()

	// init the aws client
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	// set aws region
	cfg.Region = awsRegion

	// pass aws client config
	awsClient = ec2.New(cfg)

	// start the poller
	StartPoller()
}
