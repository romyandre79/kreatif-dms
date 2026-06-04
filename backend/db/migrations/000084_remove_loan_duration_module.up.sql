-- 000084_remove_loan_duration_module.up.sql

-- 1. Remove loan_duration module from system_modules and role_permissions
DELETE FROM role_permissions WHERE module_id = 'loan_duration';
DELETE FROM system_modules WHERE id = 'loan_duration';

-- 2. Seed default_loan_duration_days setting under category general
INSERT INTO system_settings (category, key, value, value_type, description)
VALUES ('general', 'default_loan_duration_days', '7', 'integer', 'Durasi peminjaman dokumen default (hari)')
ON CONFLICT (category, key) DO UPDATE SET
    description = EXCLUDED.description;
