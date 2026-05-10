-- 000019_sync_system_modules.up.sql
-- Drop and recreate with new fields for dynamic menu
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS system_modules;

CREATE TABLE IF NOT EXISTS system_modules (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) NOT NULL,
    path VARCHAR(255),
    icon VARCHAR(50),
    allowed_actions TEXT[] NOT NULL,
    sort_order INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    module_id VARCHAR(50) NOT NULL REFERENCES system_modules(id) ON DELETE CASCADE,
    action VARCHAR(20) NOT NULL,
    PRIMARY KEY (role_id, module_id, action)
);

-- Seed updated modules with standardized actions, paths, and icons
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order) VALUES
-- Dashboard
('dashboard', 'Dashboard Utama', 'Dashboard', '/dashboard', 'LucideLayoutDashboard', ARRAY['VIEW', 'PRINT'], 1),

-- Registration
('new_doc', 'Registrasi Dokumen Baru', 'Registration', '/documents/upload', 'LucidePlusCircle', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 10),
('legacy_mig', 'Migrasi Data Legacy', 'Registration', '/registration/migration', 'LucideUploadCloud', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 11),
('staging', 'Inbound Staging', 'Registration', '/registration/staging', 'LucideInbox', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 12),

-- Warehouse
('structure', 'Struktur Gudang', 'Warehouse', '/warehouse/structure', 'LucideHome', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 20),
('occupancy', 'Okupansi Ruang', 'Warehouse', '/warehouse/occupancy', 'LucidePieChart', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 21),
('zonation', 'Zonasi Gudang', 'Warehouse', '/warehouse/zonation', 'LucideMap', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 22),
('labels', 'Cetak Label QR', 'Warehouse', '/warehouse/labels', 'LucideQrCode', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 23),

-- Search
('search', 'Pencarian Global', 'Search', '/documents', 'LucideSearch', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 30),

-- Circulation
('pickup', 'Pengambilan Dokumen', 'Circulation', '/circulation/pickup', 'LucideShoppingBag', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 40),
('checkout', 'Peminjaman Dokumen', 'Circulation', '/circulation/checkout', 'LucideLogOut', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 41),
('checkin', 'Pengembalian Dokumen', 'Circulation', '/circulation/checkin', 'LucideLogIn', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 42),
('overdue', 'Monitoring Keterlambatan', 'Circulation', '/circulation/overdue', 'LucideClock', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 43),
('softcopy', 'Permintaan Softcopy', 'Circulation', '/circulation/softcopy', 'LucideFileJson', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 44),

-- Stock Take
('audit_missions', 'Misi Audit', 'Stock Take', '/stock/missions', 'LucideClipboardList', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 50),
('scan_exec', 'Eksekusi Scanning', 'Stock Take', '/stock/scan', 'LucideScanLine', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 51),
('reconciliation', 'Rekonsiliasi Data', 'Stock Take', '/stock/reconciliation', 'LucideGitCompare', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 52),

-- Retention
('approaching', 'Mendekati Masa Retensi', 'Retention', '/retention/approaching', 'LucideHistory', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 60),
('history', 'Riwayat Pemusnahan', 'Retention', '/retention/history', 'LucideTrash2', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 61),

-- Reports
('stats', 'Statistik Sistem', 'Reports', '/reports/statistics', 'LucideBarChart3', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 70),
('capacity', 'Kapasitas Media Simpan', 'Reports', '/reports/capacity', 'LucideHardDrive', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 71),
('audit_trail', 'Jejak Audit Dokumen', 'Reports', '/reports/audit', 'LucideActivity', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 72),

-- Configuration
('company', 'Manajemen Perusahaan', 'Configuration', '/config/company', 'LucideBuilding2', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 80),
('branch', 'Manajemen Cabang', 'Configuration', '/config/branch', 'LucideMapPin', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 81),
('dept', 'Manajemen Departemen', 'Configuration', '/config/department', 'LucideUsers', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 82),
('rack', 'Manajemen Rak', 'Configuration', '/config/rack', 'LucideServer', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 83),
('box', 'Manajemen Box', 'Configuration', '/config/box', 'LucideArchive', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 84),
('ordner', 'Manajemen Ordner', 'Configuration', '/config/ordner', 'LucideFolder', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 85),
('type', 'Jenis Dokumen', 'Configuration', '/config/document-type', 'LucideFileText', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 86),
('params', 'Parameter Sistem', 'Configuration', '/config/params', 'LucideSettings2', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 87),
('integration', 'Integrasi Sistem', 'Configuration', '/config/integration', 'LucideActivity', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 88),
('audit_logs', 'Log Audit Sistem', 'Configuration', '/admin/audit-logs', 'LucideHistory', ARRAY['VIEW', 'EXPORT', 'DOWNLOAD', 'PRINT'], 89),

-- Administration
('users', 'Manajemen User', 'Administration', '/admin/users', 'LucideUserCog', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 90),
('roles', 'Maintain Role', 'Administration', '/admin/roles', 'LucideShieldAlert', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 91),
('modules', 'System Modules', 'Administration', '/admin/modules', 'LucideLayers', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 92);

-- Grant ALL permissions to superadmin
DO $$
DECLARE
    superadmin_role_id INT;
BEGIN
    -- Ensure superadmin role exists and get its ID
    SELECT id INTO superadmin_role_id FROM roles WHERE name = 'superadmin';
    
    IF superadmin_role_id IS NOT NULL THEN
        -- Clear existing permissions for superadmin
        DELETE FROM role_permissions WHERE role_id = superadmin_role_id;
        
        -- Grant all permissions for all modules
        INSERT INTO role_permissions (role_id, module_id, action)
        SELECT superadmin_role_id, m.id, a
        FROM system_modules m
        CROSS JOIN unnest(m.allowed_actions) a
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;
END $$;
