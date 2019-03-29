package backend

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

var (
	awsAccessKeyId     string
	awsAccessKeySecret string
)

// aws CredentialProvider implementation
type awsCredProvider struct{}

func (c awsCredProvider) Retrieve() (aws.Credentials, error) {
	awsCreds := aws.Credentials{
		AccessKeyID:     awsAccessKeyId,
		SecretAccessKey: awsAccessKeySecret,
		CanExpire:       false,
	}

	return awsCreds, nil
}

// Look for AWS credentials from Env
func detectEnvAPIKeys() (present bool) {
	if os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") != "" {
		present = true
	}

	return
}

// load AWS creds from default provider if exists, else custom
func loadAWSConfig() (cfg aws.Config, err error) {
	cfg, err = external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
		return
	}

	// LoadDefaultAWSConfig can load config from many places
	// but we will only support the env vars in detectEnvAPIKeys
	if !detectEnvAPIKeys() {
		// block until credentials are ready
		log.Warning("aws credentials are not found. Waiting for user interaction from GUI...")
		for {
			if awsAccessKeyId != "" && awsAccessKeySecret != "" {
				cfg.Credentials = awsCredProvider{}
				break
			}
			time.Sleep(3 * time.Second)
			// reminder every 300 seconds :)
			if time.Now().Unix() % 300 == 0 {
				log.Critical("Waiting for user to enter aws credentials from GUI...")
			}
		}
	}

	return
}
