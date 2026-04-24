-- Add manager_doc_controller role
INSERT INTO roles (name, description) VALUES 
('manager_doc_controller', 'Head of Document Control with full operational management access')
ON CONFLICT (name) DO NOTHING;

-- Seed a user for manager_doc_controller
-- Password is 'password' (bcrypt hash)
INSERT INTO users (email, password_hash, full_name, role_id, status)
SELECT 'manager_dc@kreatif.id', '$2a$10$qXhx1.rvn8kO.Cg6WqO2aO/lwyVsVQkhuaVhdW2kjPpQp/hITuzY.', 'Head of Document Control', id, 'approved'
FROM roles WHERE name = 'manager_doc_controller'
ON CONFLICT (email) DO NOTHING;
