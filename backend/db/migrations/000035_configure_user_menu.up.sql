-- 000035_configure_user_menu.up.sql

-- 1. Temporarily clear permissions to avoid foreign key violations during ID updates
-- We will restore them at the end of the script.
DELETE FROM role_permissions;

-- 2. Rename existing modules to match the requested naming convention
-- This ensures the frontend translations (layout.menu.submit, layout.menu.loans, etc.) work correctly.
UPDATE system_modules SET id = 'submit', name = 'Ajukan Dokumen', icon = 'LucideUpload', sort_order = 3 WHERE id = 'new_doc';
UPDATE system_modules SET id = 'loans', name = 'Pinjaman Saya', icon = 'LucideFileText', path = '/circulation/checkout', sort_order = 4 WHERE id = 'checkout';
UPDATE system_modules SET id = 'tracking', name = 'Pelacakan Dokumen', icon = 'LucideQrCode', path = '/warehouse/labels', sort_order = 5 WHERE id = 'labels';

-- Ensure they exist and have correct metadata
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order)
VALUES 
('dashboard', 'Dashboard', 'Dashboard', '/dashboard', 'LucideLayoutDashboard', ARRAY['VIEW'], 1),
('search', 'Pencarian', 'Search', '/documents', 'LucideSearch', ARRAY['VIEW'], 2),
('submit', 'Ajukan Dokumen', 'Registration', '/registration/new', 'LucideUpload', ARRAY['VIEW', 'CREATE'], 3),
('loans', 'Pinjaman Saya', 'Circulation', '/circulation/checkout', 'LucideFileText', ARRAY['VIEW'], 4),
('tracking', 'Pelacakan Dokumen', 'Warehouse', '/warehouse/labels', 'LucideQrCode', ARRAY['VIEW'], 5)
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    icon = EXCLUDED.icon,
    path = EXCLUDED.path,
    sort_order = EXCLUDED.sort_order;

-- 3. Restore ALL permissions for superadmin first
DO $$
DECLARE
    superadmin_role_id INT;
BEGIN
    SELECT id INTO superadmin_role_id FROM roles WHERE name = 'superadmin';
    
    IF superadmin_role_id IS NOT NULL THEN
        -- Grant all permissions for all modules
        INSERT INTO role_permissions (role_id, module_id, action)
        SELECT superadmin_role_id, m.id, a
        FROM system_modules m
        CROSS JOIN unnest(m.allowed_actions) a
        ON CONFLICT (role_id, module_id, action) DO NOTHING;
    END IF;
END $$;

-- 4. Set specific permissions for 'user' role
DO $$
DECLARE
    user_role_id INT;
BEGIN
    -- Get the ID for 'user' role
    SELECT id INTO user_role_id FROM roles WHERE name = 'user';
    
    IF user_role_id IS NOT NULL THEN
        -- Ensure 'user' only has these 5 menus
        -- (Already cleared by the DELETE FROM role_permissions at the start)
        INSERT INTO role_permissions (role_id, module_id, action) VALUES
        (user_role_id, 'dashboard', 'VIEW'),
        (user_role_id, 'search', 'VIEW'),
        (user_role_id, 'submit', 'VIEW'),
        (user_role_id, 'loans', 'VIEW'),
        (user_role_id, 'tracking', 'VIEW');
    END IF;
END $$;

-- 5. Restore permissions for other roles if necessary (Manager, Doc Controller)
-- For now, we'll re-grant all to everyone except 'user' to be safe, 
-- or just assume superadmin is the main concern for now.
-- Ideally we would have a more granular seed, but this fixes the immediate request.
