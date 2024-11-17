package minio

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
)

// UploadFile загружает файл в MinIO и возвращает временную ссылку
func (mc *Client) UploadFile(ctx context.Context, objectName, filePath, contentType string) (string, error) {
	// Загрузка файла
	_, err := mc.client.FPutObject(
		ctx, mc.bucketName, objectName, filePath, minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		return "", err
	}

	// Генерация временной ссылки
	expiry := time.Hour * 24
	presignedURL, err := mc.client.PresignedGetObject(ctx, mc.bucketName, objectName, expiry, nil)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
