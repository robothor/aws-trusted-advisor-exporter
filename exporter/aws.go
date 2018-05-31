package exporter

import (
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/support"
)

const SupportRegion = "us-east-1"

type SupportClient struct {
	reader *support.Support
}

type SupportReader interface {
	RefreshLimits()
}

func NewSupportClient() *SupportClient {
	// Using the SDKs default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Set the AWS Region that the service clients should use
	cfg.Region = endpoints.UsEast1RegionID

	return &SupportClient{
		reader: support.New(cfg),
	}
}

func (c *SupportClient) GetSupportedChecks() {
	req:=c.reader.DescribeTrustedAdvisorChecks(&support.DescribeTrustedAdvisorChecksInput{
		Language: "en"
	})
}
