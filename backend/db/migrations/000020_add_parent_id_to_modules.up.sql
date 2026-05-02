-- 000020_add_parent_id_to_modules.up.sql

-- 1. Add parent_id column
ALTER TABLE system_modules ADD COLUMN IF NOT EXISTS parent_id VARCHAR(50) REFERENCES system_modules(id) ON DELETE SET NULL;

-- 2. Create Parent Rows for existing categories
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order) VALUES
('cat_dashboard', 'Dashboard', '', NULL, 'LucideLayoutDashboard', ARRAY['VIEW'], 1),
('cat_registration', 'Registration', '', NULL, 'LucidePlusCircle', ARRAY['VIEW'], 10),
('cat_warehouse', 'Warehouse', '', NULL, 'LucideHome', ARRAY['VIEW'], 20),
('cat_search', 'Search', '', NULL, 'LucideSearch', ARRAY['VIEW'], 30),
('cat_circulation', 'Circulation', '', NULL, 'LucideRepeat', ARRAY['VIEW'], 40),
('cat_stock', 'Stock Take', '', NULL, 'LucideBox', ARRAY['VIEW'], 50),
('cat_retention', 'Retention', '', NULL, 'LucideTrash2', ARRAY['VIEW'], 60),
('cat_reports', 'Reports', '', NULL, 'LucideBarChart3', ARRAY['VIEW'], 70),
('cat_config', 'Configuration', '', NULL, 'LucideSettings', ARRAY['VIEW'], 80),
('cat_admin', 'Administration', '', NULL, 'LucideShieldAlert', ARRAY['VIEW'], 90)
ON CONFLICT (id) DO NOTHING;

-- 3. Link existing modules to new parents
UPDATE system_modules SET parent_id = 'cat_dashboard' WHERE id = 'dashboard';
UPDATE system_modules SET parent_id = 'cat_registration' WHERE category = 'Registration';
UPDATE system_modules SET parent_id = 'cat_warehouse' WHERE category = 'Warehouse';
UPDATE system_modules SET parent_id = 'cat_search' WHERE id = 'search';
UPDATE system_modules SET parent_id = 'cat_circulation' WHERE category = 'Circulation';
UPDATE system_modules SET parent_id = 'cat_stock' WHERE category = 'Stock Take';
UPDATE system_modules SET parent_id = 'cat_retention' WHERE category = 'Retention';
UPDATE system_modules SET parent_id = 'cat_reports' WHERE category = 'Reports';
UPDATE system_modules SET parent_id = 'cat_config' WHERE category = 'Configuration';
UPDATE system_modules SET parent_id = 'cat_admin' WHERE category = 'Administration';

-- 4. Clean up: Category column is now redundant (but we keep it for now just in case)
-- ALTER TABLE system_modules DROP COLUMN category;

-- 5. Ensure superadmin has permissions for the new category rows
INSERT INTO role_permissions (role_id, module_id, action)
SELECT r.id, m.id, 'VIEW'
FROM roles r
CROSS JOIN system_modules m
WHERE r.name = 'superadmin' AND m.id LIKE 'cat_%'
ON CONFLICT DO NOTHING;
