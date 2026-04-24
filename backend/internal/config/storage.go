package config

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinIO(endpoint, accessKey, secretKey string, useSSL bool, bucketName string) *minio.Client {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("MinIO connection failed: %v", err)
	}

	// Check if bucket exists, create if not
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalf("MinIO bucket check failed: %v", err)
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("MinIO bucket creation failed: %v", err)
		}
		log.Printf("MinIO bucket '%s' created\n", bucketName)
	}

	log.Println("MinIO connection established")
	return minioClient
}
