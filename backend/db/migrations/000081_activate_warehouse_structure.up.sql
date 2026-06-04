-- 000081_activate_warehouse_structure.up.sql

-- Grant VIEW permission for 'cat_warehouse' and 'structure' modules to 'manajer' and 'user' roles
DO $$
DECLARE
    target_role_id INT;
    target_role_name TEXT;
    target_roles TEXT[] := ARRAY['user', 'manajer', 'superadmin', 'admin doc controller', 'kepala doc controller'];
BEGIN
    FOR target_role_name IN SELECT unnest(target_roles)
    LOOP
        SELECT id INTO target_role_id FROM roles WHERE name = target_role_name;
        
        IF target_role_id IS NOT NULL THEN
            -- Grant VIEW permission for the parent category 'cat_warehouse'
            IF EXISTS (SELECT 1 FROM system_modules WHERE id = 'cat_warehouse') THEN
                INSERT INTO role_permissions (role_id, module_id, action)
                VALUES (target_role_id, 'cat_warehouse', 'VIEW')
                ON CONFLICT (role_id, module_id, action) DO NOTHING;
            END IF;

            -- Grant VIEW permission for the 'structure' module (Location Structure)
            IF EXISTS (SELECT 1 FROM system_modules WHERE id = 'structure') THEN
                INSERT INTO role_permissions (role_id, module_id, action)
                VALUES (target_role_id, 'structure', 'VIEW')
                ON CONFLICT (role_id, module_id, action) DO NOTHING;
            END IF;
        END IF;
    END LOOP;
END $$;
