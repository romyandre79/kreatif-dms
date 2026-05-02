-- Migration: Update S3 Node Configuration with credentials
UPDATE integration_nodes 
SET 
    endpoint = 'localhost:9000',
    config_json = '{"bucket": "kreatif-dms", "access_key": "minioadmin", "secret_key": "secretpassword", "use_ssl": false}'
WHERE service_type = 'S3';

-- If it doesn't exist, insert it
INSERT INTO integration_nodes (name, service_type, driver, endpoint, is_critical, config_json)
SELECT 'Primary S3 Storage', 'S3', 'minio', 'localhost:9000', true, '{"bucket": "kreatif-dms", "access_key": "minioadmin", "secret_key": "secretpassword", "use_ssl": false}'
WHERE NOT EXISTS (SELECT 1 FROM integration_nodes WHERE service_type = 'S3');
