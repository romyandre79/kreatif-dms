-- 000037_configure_controller_menu.up.sql

-- 1. Ensure all categories (parents) exist with correct IDs
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id) VALUES
('cat_registration', 'Document Registration', 'Core', NULL, 'LucideFileText', ARRAY['VIEW'], 10, NULL),
('cat_warehouse', 'Warehouse Management', 'Core', NULL, 'LucideHome', ARRAY['VIEW'], 20, NULL),
('cat_circulation', 'Circulation', 'Core', NULL, 'LucideRepeat', ARRAY['VIEW'], 30, NULL),
('cat_stock', 'Stock Take', 'Core', NULL, 'LucideClipboardList', ARRAY['VIEW'], 40, NULL),
('cat_retention', 'Retention & Disposal', 'Core', NULL, 'LucideTrash2', ARRAY['VIEW'], 50, NULL),
('cat_reports', 'Reports', 'System', NULL, 'LucideBarChart3', ARRAY['VIEW'], 60, NULL),
('cat_config', 'Configuration', 'System', NULL, 'LucideSettings', ARRAY['VIEW'], 70, NULL)
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    icon = EXCLUDED.icon,
    sort_order = EXCLUDED.sort_order;

-- 2. Add or Update specific modules for Controller roles
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id) VALUES
-- Registration
('inbound_pre', 'Inbound Pre-Registration', 'Registration', '/registration/pre', 'LucideInbox', ARRAY['VIEW', 'CREATE', 'EDIT'], 11, 'cat_registration'),
('submit', 'New Document Input', 'Registration', '/documents/upload', 'LucidePlusCircle', ARRAY['VIEW', 'CREATE'], 12, 'cat_registration'),
('legacy_mig', 'Legacy Data Migration', 'Registration', '/registration/migration', 'LucideUploadCloud', ARRAY['VIEW', 'CREATE'], 13, 'cat_registration'),

-- Warehouse
('structure', 'Location Structure', 'Warehouse', '/warehouse/structure', 'LucideHome', ARRAY['VIEW', 'CREATE', 'EDIT'], 21, 'cat_warehouse'),
('occupancy', 'Occupancy Map', 'Warehouse', '/warehouse/occupancy', 'LucidePieChart', ARRAY['VIEW'], 22, 'cat_warehouse'),
('zonation', 'Department Zonation', 'Warehouse', '/warehouse/zonation', 'LucideMap', ARRAY['VIEW', 'CREATE', 'EDIT'], 23, 'cat_warehouse'),
('tracking', 'Print QR Labels', 'Warehouse', '/warehouse/labels', 'LucideQrCode', ARRAY['VIEW'], 24, 'cat_warehouse'),

-- Search
('search', 'Document Search', 'Search', '/documents', 'LucideSearch', ARRAY['VIEW'], 30, NULL),

-- Circulation
('pickup', 'Pickup Preparation', 'Circulation', '/circulation/pickup', 'LucideShoppingBag', ARRAY['VIEW', 'CREATE'], 31, 'cat_circulation'),
('loans', 'Checkout (Handover)', 'Circulation', '/circulation/checkout', 'LucideLogOut', ARRAY['VIEW', 'CREATE'], 32, 'cat_circulation'),
('checkin', 'Check-in (Return)', 'Circulation', '/circulation/checkin', 'LucideLogIn', ARRAY['VIEW', 'CREATE'], 33, 'cat_circulation'),
('overdue', 'Overdue Monitor', 'Circulation', '/circulation/overdue', 'LucideClock', ARRAY['VIEW'], 34, 'cat_circulation'),
('softcopy', 'Softcopy Release', 'Circulation', '/circulation/softcopy', 'LucideFileJson', ARRAY['VIEW', 'CREATE'], 35, 'cat_circulation'),

-- Stock Take
('audit_missions', 'Audit Mission List', 'Stock Take', '/stock/missions', 'LucideClipboardList', ARRAY['VIEW', 'CREATE'], 41, 'cat_stock'),
('scan_exec', 'Scan Execution', 'Stock Take', '/stock/scan', 'LucideScanLine', ARRAY['VIEW', 'CREATE'], 42, 'cat_stock'),
('reconciliation', 'Reconciliation Report', 'Stock Take', '/stock/reconciliation', 'LucideGitCompare', ARRAY['VIEW'], 43, 'cat_stock'),

-- Retention
('approaching', 'Approaching Retention', 'Retention', '/retention/approaching', 'LucideHistory', ARRAY['VIEW'], 51, 'cat_retention'),
('history', 'Disposal History', 'Retention', '/retention/history', 'LucideTrash2', ARRAY['VIEW'], 52, 'cat_retention'),

-- Reports
('stats', 'Document Statistics', 'Reports', '/reports/statistics', 'LucideBarChart3', ARRAY['VIEW'], 61, 'cat_reports'),
('capacity', 'Warehouse Capacity', 'Reports', '/reports/capacity', 'LucideHardDrive', ARRAY['VIEW'], 62, 'cat_reports'),
('audit_trail', 'Audit Trail', 'Reports', '/reports/audit', 'LucideActivity', ARRAY['VIEW'], 63, 'cat_reports'),

-- Configuration
('company', 'Company Master', 'Configuration', '/config/company', 'LucideBuilding2', ARRAY['VIEW', 'CREATE', 'EDIT'], 71, 'cat_config'),
('dept', 'Department Master', 'Configuration', '/config/department', 'LucideUsers', ARRAY['VIEW', 'CREATE', 'EDIT'], 72, 'cat_config'),
('params', 'System Parameters', 'Configuration', '/config/params', 'LucideSettings2', ARRAY['VIEW', 'CREATE', 'EDIT'], 73, 'cat_config')

ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    icon = EXCLUDED.icon,
    path = EXCLUDED.path,
    sort_order = EXCLUDED.sort_order,
    parent_id = EXCLUDED.parent_id;

-- 3. Configure Role Permissions for 'admin doc controller' and 'kepala doc controller'
DO $$
DECLARE
    roles_list TEXT[] := ARRAY['admin doc controller', 'kepala doc controller'];
    role_name TEXT;
    role_id_val INT;
BEGIN
    FOREACH role_name IN ARRAY roles_list LOOP
        SELECT id INTO role_id_val FROM roles WHERE name = role_name;
        
        IF role_id_val IS NOT NULL THEN
            -- Clear existing permissions
            DELETE FROM role_permissions WHERE role_id = role_id_val;
            
            -- Grant VIEW permission for all modules defined above + dashboard
            INSERT INTO role_permissions (role_id, module_id, action)
            SELECT role_id_val, id, 'VIEW'
            FROM system_modules
            WHERE id IN (
                'dashboard', 'cat_registration', 'inbound_pre', 'submit', 'legacy_mig', 'staging',
                'cat_warehouse', 'structure', 'occupancy', 'zonation', 'tracking',
                'search', 'cat_circulation', 'pickup', 'loans', 'checkin', 'overdue', 'softcopy',
                'cat_stock', 'audit_missions', 'scan_exec', 'reconciliation',
                'cat_retention', 'approaching', 'history',
                'cat_reports', 'stats', 'capacity', 'audit_trail',
                'cat_config', 'company', 'dept', 'params'
            );
            
            -- Also grant specific actions if needed (e.g. CREATE for some modules)
            -- For simplicity and matching the request, we stick to VIEW as it's the primary menu filter.
        END IF;
    END LOOP;
END $$;
