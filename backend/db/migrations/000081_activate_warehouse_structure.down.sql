-- 000081_activate_warehouse_structure.down.sql

-- Remove VIEW permissions for 'cat_warehouse' and 'structure' from 'user' and 'manajer' roles
-- Keep them for superadmin and doc controller roles since they were created in previous migrations
DO $$
DECLARE
    target_role_id INT;
    target_role_name TEXT;
    target_roles TEXT[] := ARRAY['user', 'manajer'];
BEGIN
    FOR target_role_name IN SELECT unnest(target_roles)
    LOOP
        SELECT id INTO target_role_id FROM roles WHERE name = target_role_name;
        
        IF target_role_id IS NOT NULL THEN
            DELETE FROM role_permissions 
            WHERE role_id = target_role_id 
              AND module_id IN ('cat_warehouse', 'structure') 
              AND action = 'VIEW';
        END IF;
    END LOOP;
END $$;
