-- 000033_add_watermark_configuration.down.sql

DELETE FROM role_permissions WHERE module_id = 'watermark';
DELETE FROM system_modules WHERE id = 'watermark';
DELETE FROM integration_nodes WHERE service_type = 'WATERMARK';
