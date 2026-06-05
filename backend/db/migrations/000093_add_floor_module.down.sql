-- 000093_add_floor_module.down.sql

-- Remove permissions associated with 'floor'
DELETE FROM role_permissions WHERE module_id = 'floor';

-- Remove the 'floor' module
DELETE FROM system_modules WHERE id = 'floor';
