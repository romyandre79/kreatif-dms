-- Migration: Seed LDAP and OCR integration nodes
-- LDAP Configuration
UPDATE integration_nodes 
SET 
    endpoint = 'localhost:3890',
    config_json = '{"base_dn": "dc=example,dc=com", "bind_dn": "uid=admin,ou=people,dc=example,dc=com", "bind_pass": "password", "user_filter": "(&(objectClass=person)(uid=%s))", "use_tls": false}'
WHERE service_type = 'LDAP';

-- OCR Configuration
INSERT INTO integration_nodes (name, service_type, driver, endpoint, is_critical, config_json)
SELECT 'Primary OCR Engine', 'OCR', 'ocr-service', 'http://localhost:8000', true, '{"username": "admin", "password": "admin123"}'
WHERE NOT EXISTS (SELECT 1 FROM integration_nodes WHERE service_type = 'OCR');
