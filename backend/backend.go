package backend

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	// load instancetype details
	if err := loadAwsInstanceDetailsJson(); err != nil {
		log.Fatalf("could not load instance type details: %v", err)
	}
}

func StartBackendDeamon(cfgFile string) {

	// init the config
	ConfigInit(cfgFile, true)

	// start http server
	go StartServer()

	// init the aws client
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	// set aws region
	cfg.Region = awsRegion

	//defaultResolver := endpoints.NewDefaultResolver()
	myCustomResolver := func(service, region string) (aws.Endpoint, error) {
		//if service == endpoints.Ec2ServiceID {
		return aws.Endpoint{
			URL:           "http://127.0.0.1:8000/aws-mock/ec2-endpoint/",
			SigningRegion: "custom-signing-region",
		}, nil
		//}

		//return defaultResolver.ResolveEndpoint(service, region)
	}

	cfg.EndpointResolver = aws.EndpointResolverFunc(myCustomResolver)

	// pass aws client config
	awsClient = ec2.New(cfg)

	// start the poller
	StartPoller()
}
