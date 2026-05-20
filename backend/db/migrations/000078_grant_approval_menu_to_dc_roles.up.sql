-- 000078_grant_approval_menu_to_dc_roles.up.sql
-- Grant Approval Management menu (cat_approval, sub_docs, sub_loans, sub_ext)
-- to 'admin doc controller' and 'kepala doc controller' roles.

DO $$
DECLARE
    roles_list TEXT[] := ARRAY['admin doc controller', 'kepala doc controller'];
    role_name TEXT;
    role_id_val INT;
BEGIN
    FOREACH role_name IN ARRAY roles_list LOOP
        SELECT id INTO role_id_val FROM roles WHERE name = role_name;

        IF role_id_val IS NOT NULL THEN
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES
                (role_id_val, 'cat_approval', 'VIEW'),
                (role_id_val, 'sub_docs',     'VIEW'),
                (role_id_val, 'sub_loans',    'VIEW'),
                (role_id_val, 'sub_ext',      'VIEW')
            ON CONFLICT DO NOTHING;
        END IF;
    END LOOP;
END $$;
