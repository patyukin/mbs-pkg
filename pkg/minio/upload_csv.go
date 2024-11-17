package minio

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
)

func (mc *Client) UploadCSVBuffer(ctx context.Context, objectName string, buf *bytes.Buffer) (string, error) {
	if buf.Len() == 0 {
		return "", fmt.Errorf("failed to upload empty buffer")
	}

	_, err := mc.client.PutObject(
		ctx, mc.bucketName, objectName, buf, int64(buf.Len()), minio.PutObjectOptions{ContentType: "text/csv"},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload buffer to MinIO: %v", err)
	}

	expiry := time.Hour * 24
	presignedURL, err := mc.client.PresignedGetObject(ctx, mc.bucketName, objectName, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
	}

	return presignedURL.String(), nil
}
