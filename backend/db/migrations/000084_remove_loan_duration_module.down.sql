-- 000084_remove_loan_duration_module.down.sql

-- 1. Remove default_loan_duration_days setting
DELETE FROM system_settings WHERE category = 'general' AND key = 'default_loan_duration_days';

-- 2. Restore loan_duration module
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
) ON CONFLICT (id) DO NOTHING;

-- 3. Restore permissions for admin and superadmin
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
