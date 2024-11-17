package s3client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
)

func (c *Client) UploadLogReport(ctx context.Context, buf *bytes.Buffer) (string, error) {
	log.Debug().Msgf("UploadLogReport, buf: %v", buf)

	keyName := fmt.Sprintf("%s/%s.log", time.Now().Format("2006-01-02"), uuid.New().String())
	uploader := manager.NewUploader(c.client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(keyName),
		Body:   buf,
		ACL:    s3types.ObjectCannedACLPrivate,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	presignParams := &s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(keyName),
	}

	presignDuration := time.Hour * 24
	presignResult, err := c.presignClient.PresignGetObject(ctx, presignParams, func(opts *s3.PresignOptions) {
		opts.Expires = presignDuration
	})
	if err != nil {
		return "", fmt.Errorf("failed to presign file: %w", err)
	}

	return presignResult.URL, nil
}
