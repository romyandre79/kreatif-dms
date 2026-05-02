-- Migration: Force update OCR credentials
UPDATE integration_nodes 
SET 
    name = 'Primary OCR Engine',
    endpoint = 'http://127.0.0.1:8000',
    config_json = '{"username": "admin", "password": "admin123"}'
WHERE service_type = 'OCR';
