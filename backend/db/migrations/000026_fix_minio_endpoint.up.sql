-- Migration: Fix MinIO endpoint to use 127.0.0.1 instead of localhost
UPDATE integration_nodes 
SET 
    endpoint = '127.0.0.1:9000'
WHERE service_type = 'S3';
