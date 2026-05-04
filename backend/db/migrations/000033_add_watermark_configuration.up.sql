-- 000033_add_watermark_configuration.up.sql

-- 1. Insert Watermark Module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'watermark', 
    'Konfigurasi Watermark', 
    'Configuration', 
    '/config/watermark', 
    'LucideLayers', 
    ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 
    89,
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
        -- Grant all permissions for watermark module
        INSERT INTO role_permissions (role_id, module_id, action)
        SELECT superadmin_role_id, 'watermark', a
        FROM unnest(ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT']) a
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;
END $$;

-- 3. Seed Default Watermark Settings (into integration_nodes)
-- Using integration_nodes for consistency with other services
INSERT INTO integration_nodes (name, service_type, endpoint, is_active, config_json)
SELECT 
    'Default Watermark Configuration',
    'WATERMARK',
    'internal',
    true,
    '{"type": "text", "text": "CONFIDENTIAL - {user} - {date}", "opacity": 0.3, "position": "diagonal"}'
WHERE NOT EXISTS (
    SELECT 1 FROM integration_nodes WHERE service_type = 'WATERMARK'
);
