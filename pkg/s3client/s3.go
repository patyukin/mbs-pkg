package s3client

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	client        *s3.Client
	bucketName    string
	presignClient *s3.PresignClient
}

func New(ctx context.Context, bucketName string, accessKey string, secretKey string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ru-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load config, %v", err)
	}

	cl := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(cl)

	return &Client{
		client:        cl,
		bucketName:    bucketName,
		presignClient: presignClient,
	}, nil
}
