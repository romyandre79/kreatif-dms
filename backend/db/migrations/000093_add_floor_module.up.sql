-- 000093_add_floor_module.up.sql

-- Insert the 'floor' module into system_modules
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES ('floor', 'Manajemen Lantai', 'Configuration', '/config/floor', 'LucideLayers', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 82, 'cat_config')
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    icon = EXCLUDED.icon,
    path = EXCLUDED.path,
    sort_order = EXCLUDED.sort_order,
    parent_id = EXCLUDED.parent_id;

-- Grant permissions to relevant roles
DO $$
DECLARE
    roles_list TEXT[] := ARRAY['admin', 'superadmin', 'manajer', 'admin doc controller', 'kepala doc controller'];
    role_name TEXT;
    role_id_val INT;
BEGIN
    FOREACH role_name IN ARRAY roles_list LOOP
        SELECT id INTO role_id_val FROM roles WHERE name = role_name;
        
        IF role_id_val IS NOT NULL THEN
            -- Insert specific action permissions for the module
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (role_id_val, 'floor', 'VIEW'),
                   (role_id_val, 'floor', 'CREATE'),
                   (role_id_val, 'floor', 'EDIT'),
                   (role_id_val, 'floor', 'DELETE'),
                   (role_id_val, 'floor', 'EXPORT'),
                   (role_id_val, 'floor', 'DOWNLOAD'),
                   (role_id_val, 'floor', 'PRINT')
            ON CONFLICT (role_id, module_id, action) DO NOTHING;
        END IF;
    END LOOP;
END $$;
