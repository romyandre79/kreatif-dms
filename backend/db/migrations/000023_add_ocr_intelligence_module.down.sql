-- 000023_add_ocr_intelligence_module.down.sql

DELETE FROM role_permissions WHERE module_id = 'ocr';
DELETE FROM system_modules WHERE id = 'ocr';
