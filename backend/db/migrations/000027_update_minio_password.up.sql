-- Migration: Update MinIO credentials to the correct ones (minioadmin123)
UPDATE integration_nodes 
SET 
    config_json = '{"bucket": "kreatif-dms", "access_key": "minioadmin", "secret_key": "minioadmin123", "use_ssl": false}'
WHERE service_type = 'S3';
