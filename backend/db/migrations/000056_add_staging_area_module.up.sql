-- 000056_add_staging_area_module.up.sql

INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES 
('staging_area', 'Staging Area', 'Registration', '/intake/staging', 'LucideHourglass', ARRAY['VIEW', 'CREATE', 'UPDATE'], 46, 'cat_registration')
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    path = EXCLUDED.path,
    icon = EXCLUDED.icon,
    parent_id = EXCLUDED.parent_id;

-- Grant access to Admin Doc Controller, Kepala Doc Controller, and Superadmin
DO $$
DECLARE
    role_id_admin_dc INT;
    role_id_kepala_dc INT;
    role_id_superadmin INT;
BEGIN
    SELECT id INTO role_id_admin_dc FROM roles WHERE name = 'admin doc controller';
    SELECT id INTO role_id_kepala_dc FROM roles WHERE name = 'kepala doc controller';
    SELECT id INTO role_id_superadmin FROM roles WHERE name = 'superadmin';
    
    IF role_id_admin_dc IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_admin_dc, 'staging_area', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;

    IF role_id_kepala_dc IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_kepala_dc, 'staging_area', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;

    IF role_id_superadmin IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_superadmin, 'staging_area', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;
END $$;
