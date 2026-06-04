-- 000085_add_hierarchy_reporting_module.down.sql

DELETE FROM role_permissions WHERE module_id = 'admin_hierarchy';
DELETE FROM system_modules WHERE id = 'admin_hierarchy';
