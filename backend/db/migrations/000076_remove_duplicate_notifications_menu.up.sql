-- 000076_remove_duplicate_notifications_menu.up.sql

DELETE FROM role_permissions WHERE module_id = 'notifications_menu';
DELETE FROM system_modules WHERE id = 'notifications_menu';
