package config

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinIO(endpoint, accessKey, secretKey string, useSSL bool, bucketName string) (*minio.Client, error) {
	if endpoint == "" {
		return nil, nil
	}
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Printf("ERROR: MinIO client creation failed: %v\n", err)
		return nil, err
	}

	// Check if bucket exists, create if not
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Printf("WARNING: MinIO bucket check failed: %v\n", err)
		return minioClient, err
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Printf("WARNING: MinIO bucket creation failed: %v\n", err)
			return minioClient, err
		}
		log.Printf("MinIO bucket '%s' created\n", bucketName)
	}

	log.Println("MinIO connection established")
	return minioClient, nil
}
