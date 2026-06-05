-- 000089_add_dept_zonation_module.up.sql

-- 1. Create Module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'dept_zonation', 
    'Departemen Zoning Control Center', 
    'Configuration', 
    '/config/dept-zonation', 
    'LucideMap', 
    ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE'], 
    75, 
    'cat_config'
) ON CONFLICT (id) DO NOTHING;

-- 2. Add Permissions for admin, superadmin, manajer
DO $$
DECLARE
    role_rec RECORD;
    act TEXT;
BEGIN
    FOR role_rec IN SELECT id FROM roles WHERE name IN ('admin', 'superadmin', 'manajer')
    LOOP
        FOREACH act IN ARRAY ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE']
        LOOP
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (role_rec.id, 'dept_zonation', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;
