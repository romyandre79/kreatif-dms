-- 000049_grant_notifications_to_user.down.sql

DO $$
DECLARE
    user_role_id INT;
BEGIN
    SELECT id INTO user_role_id FROM roles WHERE name = 'user';
    
    IF user_role_id IS NOT NULL THEN
        DELETE FROM role_permissions 
        WHERE role_id = user_role_id AND module_id = 'notifications' AND action = 'VIEW';
    END IF;
END $$;
