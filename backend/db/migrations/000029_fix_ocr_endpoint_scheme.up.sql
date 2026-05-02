-- Migration: Ensure OCR endpoint has http scheme
UPDATE integration_nodes 
SET 
    endpoint = 'http://127.0.0.1:8000'
WHERE service_type = 'OCR';
