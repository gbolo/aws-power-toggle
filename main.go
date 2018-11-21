package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// load instancetype details
	if err := loadAwsInstanceDetailsJson(); err != nil {
		log.Fatalf("could not load instance type details: %v", err)
	}

	// init config and logging
	ConfigInit("")

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
