-- 000021_add_pin_management_module.up.sql

-- 1. Add PIN security fields to users
ALTER TABLE users ADD COLUMN IF NOT EXISTS pin_status VARCHAR(20) DEFAULT 'NOT SET';
ALTER TABLE users ADD COLUMN IF NOT EXISTS pin_failed_attempts INT DEFAULT 0;
ALTER TABLE users ADD COLUMN IF NOT EXISTS pin_locked_until TIMESTAMPTZ;
ALTER TABLE users ADD COLUMN IF NOT EXISTS pin_updated_at TIMESTAMPTZ;

-- 2. Register PIN Management module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES ('admin_pin', 'PIN Pengguna', 'Administration', '/admin/security/pin', 'LucideLock', ARRAY['VIEW', 'EDIT', 'RESET'], 93, 'cat_admin')
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    category = EXCLUDED.category,
    path = EXCLUDED.path,
    icon = EXCLUDED.icon,
    parent_id = EXCLUDED.parent_id;

-- 3. Add PIN-specific system settings
-- First, clean up any existing ones to avoid duplicates if re-run
DELETE FROM system_settings WHERE category = 'Security' AND key IN (
    'pin_approval_mandatory', 'pin_security_download_print', 'pin_archive_verification',
    'pin_length', 'pin_max_attempts', 'pin_lockout_duration', 'pin_expiry_days'
);

INSERT INTO system_settings (id, category, key, value, value_type, description) VALUES
(gen_random_uuid(), 'Security', 'pin_approval_mandatory', 'true', 'boolean', 'Wajibkan PIN untuk persetujuan dokumen'),
(gen_random_uuid(), 'Security', 'pin_security_download_print', 'true', 'boolean', 'Wajibkan PIN sebelum akses ekspor'),
(gen_random_uuid(), 'Security', 'pin_archive_verification', 'false', 'boolean', 'Wajibkan PIN untuk pengarsipan permanen'),
(gen_random_uuid(), 'Security', 'pin_length', '6', 'number', 'Panjang PIN'),
(gen_random_uuid(), 'Security', 'pin_max_attempts', '3', 'number', 'Batas percobaan PIN'),
(gen_random_uuid(), 'Security', 'pin_lockout_duration', '15', 'number', 'Durasi lockout PIN (menit)'),
(gen_random_uuid(), 'Security', 'pin_expiry_days', '90', 'number', 'Masa berlaku PIN (hari)');

-- 4. Grant permissions to superadmin
DO $$
DECLARE
    superadmin_role_id INT;
BEGIN
    SELECT id INTO superadmin_role_id FROM roles WHERE name = 'superadmin';
    IF superadmin_role_id IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES 
        (superadmin_role_id, 'admin_pin', 'VIEW'),
        (superadmin_role_id, 'admin_pin', 'EDIT'),
        (superadmin_role_id, 'admin_pin', 'RESET')
        ON CONFLICT DO NOTHING;
    END IF;
END $$;
