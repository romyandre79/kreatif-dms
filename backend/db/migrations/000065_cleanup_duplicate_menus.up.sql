-- 000065_cleanup_duplicate_menus.up.sql

-- 1. Remove duplicate Parameter Sistem (ID: params from 000019)
-- We keep ID 'config_params' from 000046 as it's more specific, 
-- but we make sure the translation key is handled.
-- Actually, let's keep 'params' as the ID because it's used in translations and other places.
-- So we delete 'config_params' and ensure 'params' has the correct config.

DELETE FROM role_permissions WHERE module_id = 'config_params';
DELETE FROM system_modules WHERE id = 'config_params';

UPDATE system_modules SET 
    name = 'Parameter Sistem',
    category = 'Configuration',
    path = '/config/params',
    icon = 'LucideSettings2',
    sort_order = 99,
    parent_id = 'cat_config'
WHERE id = 'params';

-- 2. Fix Announcements path
-- The file is at frontend/app/pages/config/announcements.vue -> /config/announcements
UPDATE system_modules SET 
    path = '/config/announcements',
    category = 'Configuration',
    parent_id = 'cat_config'
WHERE id = 'announcements';

-- 3. Ensure permissions for 'params' are correctly set for admin/superadmin
DO $$
DECLARE
    role_rec RECORD;
    act TEXT;
BEGIN
    FOR role_rec IN SELECT id FROM roles WHERE name IN ('admin', 'superadmin')
    LOOP
        FOREACH act IN ARRAY ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT']
        LOOP
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (role_rec.id, 'params', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;
