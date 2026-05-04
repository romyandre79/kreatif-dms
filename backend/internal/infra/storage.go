package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/encrypt"
	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/repository"
)

type StorageService struct {
	cfg    config.Config
	repo   repository.Querier
	client *minio.Client
	bucket string
}

func NewStorageService(cfg config.Config, repo repository.Querier, client *minio.Client, bucket string) *StorageService {
	return &StorageService{cfg: cfg, repo: repo, client: client, bucket: bucket}
}

type storageNodeConfig struct {
	Bucket    string      `json:"bucket"`
	AccessKey string      `json:"access_key"`
	SecretKey string      `json:"secret_key"`
	UseSSL    interface{} `json:"use_ssl"`
}

func (s *StorageService) getClient(ctx context.Context) (*minio.Client, string, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "S3")
	if err != nil {
		log.Printf("[StorageService] S3 node not found in DB, using default client: %v", err)
		if s.client == nil {
			return nil, "", fmt.Errorf("storage client not initialized (check MinIO configuration in .env)")
		}
		return s.client, s.bucket, nil
	}

	var nodeCfg storageNodeConfig
	if err := json.Unmarshal(node.ConfigJson, &nodeCfg); err != nil {
		return s.client, s.bucket, fmt.Errorf("failed to parse S3 config JSON: %v", err)
	}

	// Safely parse UseSSL (could be bool or string)
	useSSL := false
	switch v := nodeCfg.UseSSL.(type) {
	case bool:
		useSSL = v
	case string:
		useSSL = (v == "true" || v == "1")
	}

	// For performance, we should cache this client. 
	// But for "everything from table" and simplicity now, we re-create or use existing.
	// Optimization: Only re-create if endpoint/keys changed.
	
	// If endpoint matches current client, just return current
	// (This is simplified, ideally we track the current endpoint in the struct)
	
	endpoint := node.Endpoint
	accessKey := nodeCfg.AccessKey
	if accessKey == "" {
		accessKey = s.cfg.MinIOAccessKey
	}
	secretKey := nodeCfg.SecretKey
	if secretKey == "" {
		secretKey = s.cfg.MinIOSecretKey
	}
	bucket := nodeCfg.Bucket
	if bucket == "" {
		bucket = s.bucket
	}

	// Create a new client for this request
	log.Printf("[StorageService] Connecting to S3 at %s with AccessKey: %s", endpoint, accessKey)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to create MinIO client for %s: %v", endpoint, err)
	}

	return client, bucket, nil
}

func (s *StorageService) Upload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string, encryptFile bool) (minio.UploadInfo, error) {
	client, bucket, err := s.getClient(ctx)
	if err != nil {
		return minio.UploadInfo{}, err
	}

	log.Printf("[StorageService] Uploading object: %s (Size: %d, Type: %s) to bucket %s", objectName, size, contentType, bucket)
	opts := minio.PutObjectOptions{
		ContentType: contentType,
	}

	if encryptFile {
		// Using SSE-S3 (Managed by MinIO)
		opts.ServerSideEncryption = encrypt.NewSSE()
	}

	info, err := client.PutObject(ctx, bucket, objectName, reader, size, opts)
	if err != nil {
		log.Printf("[StorageService] Error uploading object %s: %v", objectName, err)
		return info, err
	}

	log.Printf("[StorageService] Object uploaded successfully: %s", objectName)
	return info, nil
}

func (s *StorageService) Download(ctx context.Context, objectName string) (io.ReadCloser, error) {
	client, bucket, err := s.getClient(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[StorageService] Downloading object: %s from bucket %s", objectName, bucket)
	return client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
}

func (s *StorageService) Delete(ctx context.Context, objectName string) error {
	client, bucket, err := s.getClient(ctx)
	if err != nil {
		return err
	}

	log.Printf("[StorageService] Deleting object: %s from bucket %s", objectName, bucket)
	err = client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Printf("[StorageService] Error deleting object %s: %v", objectName, err)
	}
	return err
}
func (s *StorageService) TestConnection(ctx context.Context, endpoint, accessKey, secretKey, bucket string, useSSL bool) error {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("failed to create MinIO client: %v", err)
	}

	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %v", err)
	}
	if !exists {
		return fmt.Errorf("bucket '%s' does not exist", bucket)
	}

	return nil
}
