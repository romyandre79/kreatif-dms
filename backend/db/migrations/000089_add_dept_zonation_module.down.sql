-- 000089_add_dept_zonation_module.down.sql
DELETE FROM role_permissions WHERE module_id = 'dept_zonation';
DELETE FROM system_modules WHERE id = 'dept_zonation';
