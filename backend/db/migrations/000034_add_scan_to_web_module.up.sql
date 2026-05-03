-- 000034_add_scan_to_web_module.up.sql

-- 1. Insert Scan to Web Module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'scan_to_web', 
    'Scan to Web Launcher', 
    'Configuration', 
    '/config/scan-to-web', 
    'LucidePrinter', 
    ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE'], 
    90,
    'cat_config'
)
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    category = EXCLUDED.category,
    path = EXCLUDED.path,
    icon = EXCLUDED.icon,
    parent_id = EXCLUDED.parent_id;

-- 2. Grant Permissions to superadmin
DO $$
DECLARE
    superadmin_role_id INT;
BEGIN
    SELECT id INTO superadmin_role_id FROM roles WHERE name = 'superadmin';
    
    IF superadmin_role_id IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES 
            (superadmin_role_id, 'scan_to_web', 'VIEW'),
            (superadmin_role_id, 'scan_to_web', 'CREATE'),
            (superadmin_role_id, 'scan_to_web', 'EDIT'),
            (superadmin_role_id, 'scan_to_web', 'DELETE')
        ON CONFLICT DO NOTHING;
    END IF;
END $$;
