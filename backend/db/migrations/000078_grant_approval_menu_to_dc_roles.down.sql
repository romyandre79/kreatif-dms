-- 000078_grant_approval_menu_to_dc_roles.down.sql

DO $$
DECLARE
    roles_list TEXT[] := ARRAY['admin doc controller', 'kepala doc controller'];
    role_name TEXT;
    role_id_val INT;
BEGIN
    FOREACH role_name IN ARRAY roles_list LOOP
        SELECT id INTO role_id_val FROM roles WHERE name = role_name;

        IF role_id_val IS NOT NULL THEN
            DELETE FROM role_permissions
            WHERE role_id = role_id_val
              AND module_id IN ('cat_approval', 'sub_docs', 'sub_loans', 'sub_ext')
              AND action = 'VIEW';
        END IF;
    END LOOP;
END $$;
