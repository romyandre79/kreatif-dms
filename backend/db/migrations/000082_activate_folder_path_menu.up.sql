-- 000082_activate_folder_path_menu.up.sql

-- Grant VIEW and EDIT permissions for the 'config_folder_path' module to admin and doc controller roles
DO $$
DECLARE
    target_role_id INT;
    target_role_name TEXT;
    target_roles TEXT[] := ARRAY['superadmin', 'admin doc controller', 'kepala doc controller'];
    act TEXT;
BEGIN
    FOR target_role_name IN SELECT unnest(target_roles)
    LOOP
        SELECT id INTO target_role_id FROM roles WHERE name = target_role_name;
        
        IF target_role_id IS NOT NULL THEN
            FOREACH act IN ARRAY ARRAY['VIEW', 'EDIT']
            LOOP
                IF EXISTS (SELECT 1 FROM system_modules WHERE id = 'config_folder_path') THEN
                    INSERT INTO role_permissions (role_id, module_id, action)
                    VALUES (target_role_id, 'config_folder_path', act)
                    ON CONFLICT DO NOTHING;
                END IF;
            END LOOP;
        END IF;
    END LOOP;
END $$;
