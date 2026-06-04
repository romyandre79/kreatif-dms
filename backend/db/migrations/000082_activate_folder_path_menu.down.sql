-- 000082_activate_folder_path_menu.down.sql

-- Remove permissions for 'config_folder_path' from doc controller roles
-- Keep them for superadmin role since they were created in the original migration (000072)
DO $$
DECLARE
    target_role_id INT;
    target_role_name TEXT;
    target_roles TEXT[] := ARRAY['admin doc controller', 'kepala doc controller'];
BEGIN
    FOR target_role_name IN SELECT unnest(target_roles)
    LOOP
        SELECT id INTO target_role_id FROM roles WHERE name = target_role_name;
        
        IF target_role_id IS NOT NULL THEN
            DELETE FROM role_permissions 
            WHERE role_id = target_role_id 
              AND module_id = 'config_folder_path';
        END IF;
    END LOOP;
END $$;
