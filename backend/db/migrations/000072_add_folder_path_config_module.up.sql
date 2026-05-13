-- 000072_add_folder_path_config_module.up.sql

-- 1. Create Module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'config_folder_path', 
    'Folder Path', 
    'Configuration', 
    '/config/folder-path', 
    'LucideFolders', 
    ARRAY['VIEW', 'EDIT'], 
    87, 
    'cat_config'
) ON CONFLICT (id) DO NOTHING;

-- 2. Add Permissions for admin and superadmin
DO $$
DECLARE
    role_rec RECORD;
    act TEXT;
BEGIN
    FOR role_rec IN SELECT id FROM roles WHERE name IN ('admin', 'superadmin')
    LOOP
        FOREACH act IN ARRAY ARRAY['VIEW', 'EDIT']
        LOOP
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (role_rec.id, 'config_folder_path', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;

-- 3. Seed initial folder path setting
INSERT INTO system_settings (category, key, value, value_type, description)
VALUES 
    ('explorer', 'folder_path', 'company,branch,department,year,rack,box,ordner', 'string', 'Urutan hierarki folder di Document Explorer'),
    ('explorer', 'create_missing_folders', 'true', 'boolean', 'Otomatis buat folder jika tidak ada'),
    ('explorer', 'apply_retention_tag', 'false', 'boolean', 'Terapkan tag retensi otomatis')
ON CONFLICT (category, key) DO UPDATE SET 
    description = EXCLUDED.description;
