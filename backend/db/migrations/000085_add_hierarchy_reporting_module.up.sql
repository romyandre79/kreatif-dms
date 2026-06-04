-- 000085_add_hierarchy_reporting_module.up.sql

-- 1. Create Module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'admin_hierarchy', 
    'Struktur & Hierarki', 
    'Administration', 
    '/admin/hierarchy', 
    'LucideGitCompare', 
    ARRAY['VIEW', 'EDIT'], 
    94, 
    'cat_admin'
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
            VALUES (role_rec.id, 'admin_hierarchy', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;
