package backend

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	// load instancetype details
	if err := loadAwsInstanceDetailsJSON(); err != nil {
		log.Fatalf("could not load instance type details: %v", err)
	}
}

// StartBackendDeamon Blocking function that starts the backend process
func StartBackendDeamon(cfgFile string) {

	// init the config
	ConfigInit(cfgFile, true)

	// start http server
	go startHTTPServer()

	// init the aws clients
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	awsClients = make(map[string]*ec2.Client, len(awsRegions))
	awsASGClients = make(map[string]*autoscaling.Client, len(awsRegions))
	for _, region := range awsRegions {
		if region != "" {
			cfg.Region = region
			awsClients[region] = ec2.New(cfg)
			awsASGClients[region] = autoscaling.New(cfg)
		}
	}

	// start the poller
	StartPoller()
}

// TODO: for mocking the actual AWS API we can try this: https://github.com/spulec/moto or https://github.com/treelogic-swe/aws-mock

//mockAwsResolver := func(service, region string) (aws.Endpoint, error) {
//	return aws.Endpoint{
//		URL:           "http://127.0.0.1:8000/aws-mock/ec2-endpoint/",
//		SigningRegion: "custom-signing-region",
//	}, nil
//}
//
//cfg.EndpointResolver = aws.EndpointResolverFunc(mockAwsResolver)
