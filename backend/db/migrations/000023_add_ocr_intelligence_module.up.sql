-- 000023_add_ocr_intelligence_module.up.sql

-- 1. Add OCR module to Configuration category
INSERT INTO system_modules (id, name, category, parent_id, path, icon, allowed_actions, sort_order) VALUES
('ocr', 'OCR Intelligence', 'Configuration', 'cat_config', '/config/ocr', 'LucideScanLine', ARRAY['VIEW'], 65)
ON CONFLICT (id) DO NOTHING;

-- 2. Grant permission to superadmin
INSERT INTO role_permissions (role_id, module_id, action)
SELECT r.id, 'ocr', 'VIEW'
FROM roles r
WHERE r.name = 'superadmin'
ON CONFLICT DO NOTHING;
