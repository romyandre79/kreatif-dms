-- 000088_add_override_rack_module.up.sql

-- 1. Create Module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'warehouse_rack_override', 
    'Override Rak', 
    'Warehouse', 
    '/warehouse/racks/override', 
    'LucideLock', 
    ARRAY['VIEW', 'EDIT'], 
    16, 
    'cat_warehouse'
) ON CONFLICT (id) DO NOTHING;

-- 2. Add Permissions for admin and superadmin and warehouse admins
DO $$
DECLARE
    role_rec RECORD;
    act TEXT;
BEGIN
    FOR role_rec IN SELECT id FROM roles WHERE name IN ('admin', 'superadmin', 'manajer', 'user')
    LOOP
        FOREACH act IN ARRAY ARRAY['VIEW', 'EDIT']
        LOOP
            INSERT INTO role_permissions (role_id, module_id, action)
            VALUES (role_rec.id, 'warehouse_rack_override', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;
