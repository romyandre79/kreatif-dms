package infra

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/encrypt"
)

type StorageService struct {
	client *minio.Client
	bucket string
}

func NewStorageService(client *minio.Client, bucket string) *StorageService {
	return &StorageService{client: client, bucket: bucket}
}

func (s *StorageService) Upload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string, encryptFile bool) (minio.UploadInfo, error) {
	log.Printf("[StorageService] Uploading object: %s (Size: %d, Type: %s)", objectName, size, contentType)
	opts := minio.PutObjectOptions{
		ContentType: contentType,
	}

	if encryptFile {
		// Using SSE-S3 (Managed by MinIO)
		opts.ServerSideEncryption = encrypt.NewSSE()
	}

	info, err := s.client.PutObject(ctx, s.bucket, objectName, reader, size, opts)
	if err != nil {
		log.Printf("[StorageService] Error uploading object %s: %v", objectName, err)
		return info, err
	}

	log.Printf("[StorageService] Object uploaded successfully: %s", objectName)
	return info, nil
}

func (s *StorageService) Download(ctx context.Context, objectName string) (io.ReadCloser, error) {
	log.Printf("[StorageService] Downloading object: %s", objectName)
	return s.client.GetObject(ctx, s.bucket, objectName, minio.GetObjectOptions{})
}

func (s *StorageService) Delete(ctx context.Context, objectName string) error {
	log.Printf("[StorageService] Deleting object: %s", objectName)
	err := s.client.RemoveObject(ctx, s.bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Printf("[StorageService] Error deleting object %s: %v", objectName, err)
	}
	return err
}
