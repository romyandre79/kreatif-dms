-- 000051_add_inbound_registration_module.up.sql

INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order)
VALUES 
('inbound_registration', 'Penerimaan Fisik', 'Registration', '/intake/inbound', 'LucideScanLine', ARRAY['VIEW', 'CREATE', 'UPDATE'], 43)
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    path = EXCLUDED.path,
    icon = EXCLUDED.icon;

-- Grant access to Admin Doc Controller and Superadmin
DO $$
DECLARE
    role_id_admin_dc INT;
    role_id_superadmin INT;
BEGIN
    SELECT id INTO role_id_admin_dc FROM roles WHERE name = 'admin doc controller';
    SELECT id INTO role_id_superadmin FROM roles WHERE name = 'superadmin';
    
    IF role_id_admin_dc IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_admin_dc, 'inbound_registration', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;

    IF role_id_superadmin IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_superadmin, 'inbound_registration', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;
END $$;
