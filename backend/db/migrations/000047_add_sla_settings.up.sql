-- 000047_add_sla_settings.up.sql

INSERT INTO system_settings (category, key, value, value_type, description)
VALUES 
    ('workflow', 'sla_approval_hours', '24', 'integer', 'Batas waktu SLA untuk persetujuan dokumen (dalam jam)')
ON CONFLICT (category, key) DO NOTHING;
