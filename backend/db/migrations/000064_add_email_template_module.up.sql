-- 000064_add_email_template_module.up.sql

-- 1. Add module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES (
    'email_templates', 
    'Email Templates', 
    'Configuration', 
    '/config/email-templates', 
    'LucideFileJson', 
    ARRAY['VIEW', 'EDIT'], 
    95, 
    'cat_config'
) ON CONFLICT (id) DO NOTHING;

-- 2. Grant permission to Admin and Superadmin roles
-- Assuming role IDs for admin and superadmin are already known or we use subqueries
DO $$
DECLARE
    admin_role_id INT;
    superadmin_role_id INT;
BEGIN
    SELECT id INTO admin_role_id FROM roles WHERE name = 'admin';
    SELECT id INTO superadmin_role_id FROM roles WHERE name = 'superadmin';

    IF admin_role_id IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES 
            (admin_role_id, 'email_templates', 'VIEW'),
            (admin_role_id, 'email_templates', 'EDIT')
        ON CONFLICT DO NOTHING;
    END IF;

    IF superadmin_role_id IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES 
            (superadmin_role_id, 'email_templates', 'VIEW'),
            (superadmin_role_id, 'email_templates', 'EDIT')
        ON CONFLICT DO NOTHING;
    END IF;
END $$;
