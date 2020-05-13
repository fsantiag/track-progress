package configuration

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/backend/src/util"
)

// NewSession create a new connection to SQS
func NewSession() *sqs.SQS {
	session := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(endpoints.UsWest2RegionID),
		Endpoint: aws.String(util.Getenv("SQS_HOST", "http://localhost:4576")),
		Credentials: credentials.NewStaticCredentials(
			util.Getenv("SQS_ID", "id"),
			util.Getenv("SQS_SECRET", "secret"),
			util.Getenv("SQS_TOKEN", "token"),
		),
	}))

	return sqs.New(session)
}
