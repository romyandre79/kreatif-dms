-- 000074_add_notifications_module.up.sql

-- 1. Add Notifications module
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order)
VALUES (
    'notifications', 
    'Notifikasi', 
    'Administration', 
    '/notifications', 
    'LucideBell', 
    ARRAY['VIEW', 'DELETE'], 
    99
) ON CONFLICT (id) DO NOTHING;

-- 2. Grant VIEW permission to all roles
INSERT INTO role_permissions (role_id, module_id, action)
SELECT r.id, 'notifications', 'VIEW'
FROM roles r
ON CONFLICT DO NOTHING;
