-- 000050_grant_notifications_to_user.up.sql

DO $$
DECLARE
    user_role_id INT;
BEGIN
    -- Ensure the module exists first
    INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order)
    VALUES ('notifications', 'Notifications', 'System', '/notifications', 'LucideBell', ARRAY['VIEW'], 100)
    ON CONFLICT (id) DO UPDATE SET 
        path = EXCLUDED.path,
        icon = EXCLUDED.icon;

    -- Get the ID for 'user' role
    SELECT id INTO user_role_id FROM roles WHERE name = 'user';
    
    IF user_role_id IS NOT NULL THEN
        INSERT INTO role_permissions (role_id, module_id, action)
        VALUES (user_role_id, 'notifications', 'VIEW')
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;
END $$;
