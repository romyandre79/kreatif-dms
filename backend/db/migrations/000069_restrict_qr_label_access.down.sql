-- 000069_restrict_qr_label_access.down.sql

-- Restore more permissive access for the QR Label module ('tracking')
-- Including 'user' and legacy roles if they exist.

DO $$
DECLARE
    target_role_id INT;
    target_role_name TEXT;
    target_roles TEXT[] := ARRAY['superadmin', 'admin', 'admin doc controller', 'kepala doc controller', 'doc controller', 'user'];
BEGIN
    FOR target_role_name IN SELECT unnest(target_roles)
    LOOP
        SELECT id INTO target_role_id FROM roles WHERE name = target_role_name;
        
        IF target_role_id IS NOT NULL THEN
            -- Restore VIEW permission for the 'tracking' module
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (target_role_id, 'tracking', 'VIEW')
            ON CONFLICT DO NOTHING;

            -- Ensure parent category 'cat_warehouse' is visible
            IF EXISTS (SELECT 1 FROM system_modules WHERE id = 'cat_warehouse') THEN
                INSERT INTO role_permissions (role_id, module_id, action)
                VALUES (target_role_id, 'cat_warehouse', 'VIEW')
                ON CONFLICT DO NOTHING;
            END IF;
        END IF;
    END LOOP;
END $$;
