-- 000072_add_folder_path_config_module.down.sql

DELETE FROM role_permissions WHERE module_id = 'config_folder_path';
DELETE FROM system_modules WHERE id = 'config_folder_path';
DELETE FROM system_settings WHERE category = 'explorer' AND key = 'folder_path';
