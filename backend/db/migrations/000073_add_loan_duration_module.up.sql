-- 000073_add_loan_duration_module.up.sql

-- 1. Add Module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'loan_duration', 
    'Durasi Pinjam', 
    'Configuration', 
    '/config/loan-duration', 
    'LucideClock', 
    ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 
    88, 
    'cat_config'
)
ON CONFLICT (id) DO UPDATE SET 
    parent_id = 'cat_config',
    name = 'Durasi Pinjam';

-- 2. Grant Permissions to admin and superadmin
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
            VALUES (role_rec.id, 'loan_duration', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;
