package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	client     *minio.Client
	bucketName string
}

func New(ctx context.Context, endpoint, accessKey, secretKey, bucketName string, useSSL bool) (*Client, error) {
	client, err := minio.New(
		endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
			Secure: useSSL,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	minioClient := &Client{
		client:     client,
		bucketName: bucketName,
	}

	err = minioClient.ensureBucket(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure bucket: %v", err)
	}

	return minioClient, nil
}

func (mc *Client) ensureBucket(ctx context.Context) error {
	exists, err := mc.client.BucketExists(ctx, mc.bucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %v", err)
	}

	if !exists {
		err = mc.client.MakeBucket(ctx, mc.bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
	}

	return nil
}
