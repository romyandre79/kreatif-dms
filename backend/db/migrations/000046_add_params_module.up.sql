-- Add Module
INSERT INTO system_modules (id, name, parent_id, path, icon, sort_order, category)
VALUES ('config_params', 'Parameter Sistem', 'Configuration', '/config/params', 'LucideSettings2', 99, 'admin')
ON CONFLICT (id) DO NOTHING;

-- Add Permissions for admin
INSERT INTO role_permissions (role_id, module_id, action)
VALUES 
    ('admin', 'config_params', 'read'),
    ('admin', 'config_params', 'update'),
    ('superadmin', 'config_params', 'read'),
    ('superadmin', 'config_params', 'update')
ON CONFLICT DO NOTHING;

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
