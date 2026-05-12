-- Migration: Add Document Category module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES ('doc_category', 'Kategori Dokumen', 'Configuration', '/config/document-category', 'LucideLayers', ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT'], 85, 'cat_config')
ON CONFLICT (id) DO NOTHING;

-- Grant permissions to superadmin
INSERT INTO role_permissions (role_id, module_id, action)
SELECT r.id, 'doc_category', a
FROM roles r
CROSS JOIN unnest(ARRAY['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT']) a
WHERE r.name = 'superadmin'
ON CONFLICT DO NOTHING;
