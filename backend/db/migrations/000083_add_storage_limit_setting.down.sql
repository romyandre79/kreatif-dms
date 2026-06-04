-- 000083_add_storage_limit_setting.down.sql
DELETE FROM system_settings WHERE category = 'storage' AND key = 'storage_limit';
