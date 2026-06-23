-- 000091_remove_old_warehouse_zonation.down.sql

-- 1. Restore the old zonation module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'zonation', 
    'Zonasi Departemen', 
    'Warehouse', 
    '/warehouse/zonation', 
    'LucideMap', 
    ARRAY['VIEW', 'CREATE', 'EDIT'], 
    23, 
    'cat_warehouse'
) ON CONFLICT (id) DO NOTHING;

-- 2. Restore permissions for roles
DO $$
DECLARE
    role_rec RECORD;
    act TEXT;
BEGIN
    FOR role_rec IN SELECT id FROM roles WHERE name IN ('admin', 'superadmin', 'manajer', 'controller')
    LOOP
        FOREACH act IN ARRAY ARRAY['VIEW', 'CREATE', 'EDIT']
        LOOP
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (role_rec.id, 'zonation', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;
