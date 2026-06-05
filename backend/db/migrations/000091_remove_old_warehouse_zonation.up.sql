-- 000091_remove_old_warehouse_zonation.up.sql

-- 1. Remove permissions for the old zonation module
DELETE FROM role_permissions WHERE module_id = 'zonation';

-- 2. Remove the old zonation module from system_modules
DELETE FROM system_modules WHERE id = 'zonation';
