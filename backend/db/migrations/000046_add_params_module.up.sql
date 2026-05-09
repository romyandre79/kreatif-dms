-- Fix 000046_add_params_module.up.sql
-- Add Module with all required columns
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'config_params', 
    'Parameter Sistem', 
    'Configuration', 
    '/config/params', 
    'LucideSettings2', 
    ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 
    99, 
    'cat_config'
)
ON CONFLICT (id) DO UPDATE SET 
    parent_id = 'cat_config',
    allowed_actions = ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'];

-- Add Permissions for admin and superadmin
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
            VALUES (role_rec.id, 'config_params', act)
            ON CONFLICT DO NOTHING;
        END LOOP;
    END LOOP;
END $$;

-- Seed initial parameters if not exists
INSERT INTO system_settings (category, key, value, value_type, description)
VALUES 
    ('general', 'app_version', 'V3.1.2', 'string', 'Versi sistem yang ditampilkan di manifest'),
    ('general', 'app_name', 'Kreatif DMS', 'string', 'Nama aplikasi utama'),
    ('general', 'company_name', 'PT. KREATIF DIGITAL SOLUSI', 'string', 'Nama perusahaan default'),
    ('general', 'manifest_items_per_page', '15', 'integer', 'Jumlah baris dokumen per halaman pada cetak manifest'),
    ('storage', 'encryption_enabled', 'true', 'boolean', 'Aktifkan enkripsi AES-256 (SSE-S3) untuk file di MinIO'),
    ('ocr', 'auto_process', 'true', 'boolean', 'Proses OCR otomatis setelah unggah'),
    ('storage', 'max_upload_size', '100', 'integer', 'Batas maksimal unggah file (MB)')
ON CONFLICT (category, key) DO NOTHING;
