-- 000083_add_storage_limit_setting.up.sql
INSERT INTO system_settings (category, key, value, value_type, description)
VALUES ('storage', 'storage_limit', '100', 'integer', 'Batas kapasitas penyimpanan total dalam GB')
ON CONFLICT (category, key) DO NOTHING;
