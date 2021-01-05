package reports

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (r *Reports) s3Object(file string) *s3.GetObjectOutput {
	cfg := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		r.logger.Fatal().Err(err).Msg("failed to create a new S3 session")
	}

	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_UPLOAD_BUCKET")),
		Key:    aws.String(strings.Replace(file, "+", " ", -1)),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		r.logger.Error().Err(err).Msg("failed to read object")
	}

	return result
}
