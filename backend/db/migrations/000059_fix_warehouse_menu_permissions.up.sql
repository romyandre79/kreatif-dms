-- 000059_fix_warehouse_menu_permissions.up.sql

-- Ensure the Warehouse category and Print QR Label module are correctly configured
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id) VALUES
('cat_warehouse', 'Warehouse Management', 'Core', NULL, 'LucideHome', ARRAY['VIEW'], 20, NULL),
('tracking', 'Print QR Label', 'Warehouse', '/warehouse/labels', 'LucideQrCode', ARRAY['VIEW'], 24, 'cat_warehouse')
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    icon = EXCLUDED.icon,
    path = EXCLUDED.path,
    parent_id = EXCLUDED.parent_id,
    sort_order = EXCLUDED.sort_order;

-- Grant VIEW permission to key roles
DO $$
DECLARE
    role_record RECORD;
    target_roles TEXT[] := ARRAY['superadmin', 'admin', 'admin doc controller', 'kepala doc controller', 'doc controller'];
BEGIN
    FOR role_record IN SELECT id, name FROM roles WHERE name = ANY(target_roles)
    LOOP
        -- Grant to parent category
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_record.id, 'cat_warehouse', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;

        -- Grant to the module itself
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_record.id, 'tracking', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END LOOP;
END $$;
