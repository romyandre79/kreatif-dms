-- 000018_add_system_permissions.up.sql
CREATE TABLE IF NOT EXISTS system_modules (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) NOT NULL,
    allowed_actions TEXT[] NOT NULL -- Array of actions like ['VIEW', 'EDIT', etc]
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    module_id VARCHAR(50) NOT NULL REFERENCES system_modules(id) ON DELETE CASCADE,
    action VARCHAR(20) NOT NULL,
    PRIMARY KEY (role_id, module_id, action)
);

-- Seed Initial Modules
INSERT INTO system_modules (id, name, category, allowed_actions) VALUES
('global_search', 'Pencarian Global', 'Search & Navigation', ARRAY['VIEW', 'PRE', 'DL', 'PRN']),
('folder_nav', 'Navigasi Folder', 'Search & Navigation', ARRAY['VIEW', 'PRE', 'DL', 'PRN', 'UP', 'ED']),
('task_queue', 'Antrian Tugas', 'Workflow & Approval', ARRAY['VIEW', 'PRE', 'DL', 'PRN', 'APP']),
('user_management', 'Manajemen User', 'Administration', ARRAY['VIEW', 'UP', 'ED', 'DEL'])
ON CONFLICT (id) DO NOTHING;
