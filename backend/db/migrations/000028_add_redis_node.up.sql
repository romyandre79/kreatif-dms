-- Migration: Add Redis integration node
INSERT INTO integration_nodes (name, service_type, driver, endpoint, is_critical, config_json)
SELECT 'Primary Redis Cache', 'REDIS', 'redis', '127.0.0.1:6379', true, '{"db": 0, "password": ""}'
WHERE NOT EXISTS (SELECT 1 FROM integration_nodes WHERE service_type = 'REDIS');
