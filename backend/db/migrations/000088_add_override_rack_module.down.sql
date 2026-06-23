-- 000088_add_override_rack_module.down.sql

DELETE FROM role_permissions WHERE module_id = 'warehouse_rack_override';
DELETE FROM system_modules WHERE id = 'warehouse_rack_override';
