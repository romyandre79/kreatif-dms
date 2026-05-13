-- 000069_restrict_qr_label_access.up.sql

-- 1. Remove all existing permissions for the QR Label module ('tracking' and legacy 'labels')
DELETE FROM role_permissions WHERE module_id IN ('tracking', 'labels');

-- 2. Grant VIEW permission ONLY to the requested roles: superuser, admin dc, kepala dc
-- Mapping:
-- superuser -> superadmin
-- admin dc -> admin doc controller
-- kepala dc -> kepala doc controller

DO $$
DECLARE
    target_role_id INT;
    target_role_name TEXT;
    -- Note: 'admin' and 'doc controller' are explicitly excluded to follow the "ONLY" requirement.
    target_roles TEXT[] := ARRAY['superadmin', 'admin doc controller', 'kepala doc controller'];
BEGIN
    FOR target_role_name IN SELECT unnest(target_roles)
    LOOP
        SELECT id INTO target_role_id FROM roles WHERE name = target_role_name;
        
        IF target_role_id IS NOT NULL THEN
            -- Grant VIEW permission for the 'tracking' module (Cetak Label QR)
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (target_role_id, 'tracking', 'VIEW')
            ON CONFLICT DO NOTHING;

            -- Also ensure the parent category 'cat_warehouse' is visible to these roles
            IF EXISTS (SELECT 1 FROM system_modules WHERE id = 'cat_warehouse') THEN
                INSERT INTO role_permissions (role_id, module_id, action)
                VALUES (target_role_id, 'cat_warehouse', 'VIEW')
                ON CONFLICT DO NOTHING;
            END IF;
        END IF;
    END LOOP;
END $$;
