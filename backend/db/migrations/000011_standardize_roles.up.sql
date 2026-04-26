-- Standardize Roles
-- 1. Rename existing roles to match user request
UPDATE roles SET name = 'superadmin' WHERE name = 'admin';
UPDATE roles SET name = 'manajer' WHERE name = 'manager';
UPDATE roles SET name = 'admin doc controller' WHERE name = 'doc_controller';
UPDATE roles SET name = 'kepala doc controller' WHERE name = 'manager_doc_controller';

-- 2. Insert missing roles if they don't exist
INSERT INTO roles (name, description) VALUES 
('superadmin', 'Full system access'),
('manajer', 'Department manager with approval authority'),
('admin doc controller', 'Operational document management and indexing'),
('kepala doc controller', 'Strategic document management and policy oversight'),
('user', 'Standard staff access for document interaction')
ON CONFLICT (name) DO NOTHING;
