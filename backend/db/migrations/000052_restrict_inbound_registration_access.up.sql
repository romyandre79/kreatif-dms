-- 000052_restrict_inbound_registration_access.up.sql

DO $$
DECLARE
    role_id_admin_dc INT;
    role_id_kepala_dc INT;
    role_id_superadmin INT;
BEGIN
    -- Get Role IDs
    SELECT id INTO role_id_admin_dc FROM roles WHERE name = 'admin doc controller';
    SELECT id INTO role_id_kepala_dc FROM roles WHERE name = 'kepala doc controller';
    SELECT id INTO role_id_superadmin FROM roles WHERE name = 'superadmin';
    
    -- 1. Clear existing permissions for this module to be safe
    DELETE FROM role_permissions WHERE module_id = 'inbound_registration';

    -- 2. Grant VIEW permission to allowed roles
    IF role_id_admin_dc IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_admin_dc, 'inbound_registration', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;

    IF role_id_kepala_dc IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_kepala_dc, 'inbound_registration', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;

    IF role_id_superadmin IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (role_id_superadmin, 'inbound_registration', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;
END $$;
