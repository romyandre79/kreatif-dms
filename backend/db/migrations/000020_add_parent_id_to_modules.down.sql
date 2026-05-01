-- 000020_add_parent_id_to_modules.down.sql

DELETE FROM role_permissions WHERE module_id LIKE 'cat_%';
DELETE FROM system_modules WHERE id LIKE 'cat_%';
ALTER TABLE system_modules DROP COLUMN IF EXISTS parent_id;
