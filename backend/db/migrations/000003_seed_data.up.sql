-- Seed Organizations
INSERT INTO companies (id, name, address) VALUES 
('550e8400-e29b-41d4-a716-446655440000', 'Kreatif Holding', 'Jl. Kreatif No. 1, Jakarta')
ON CONFLICT (id) DO NOTHING;

INSERT INTO branches (id, company_id, name, location) VALUES 
('550e8400-e29b-41d4-a716-446655440001', '550e8400-e29b-41d4-a716-446655440000', 'Kreatif Jakarta', 'Jakarta Selatan')
ON CONFLICT (id) DO NOTHING;

INSERT INTO departments (id, branch_id, name) VALUES 
('550e8400-e29b-41d4-a716-446655440002', '550e8400-e29b-41d4-a716-446655440001', 'IT Department'),
('550e8400-e29b-41d4-a716-446655440003', '550e8400-e29b-41d4-a716-446655440001', 'Finance Department')
ON CONFLICT (id) DO NOTHING;

-- Seed Storage
INSERT INTO racks (id, department_id, name, location_detail) VALUES 
('550e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440002', 'Rack IT-A1', 'Lantai 2, Pojok Kanan')
ON CONFLICT (id) DO NOTHING;

INSERT INTO boxes (id, rack_id, name) VALUES 
('550e8400-e29b-41d4-a716-446655440005', '550e8400-e29b-41d4-a716-446655440004', 'Box IT-2024-001')
ON CONFLICT (id) DO NOTHING;

INSERT INTO ordners (id, box_id, name) VALUES 
('550e8400-e29b-41d4-a716-446655440006', '550e8400-e29b-41d4-a716-446655440005', 'Ordner Server Maintenance')
ON CONFLICT (id) DO NOTHING;

-- Seed Users (Password: password123)
-- Hash generated for 'password123'
INSERT INTO users (email, password_hash, full_name, role_id, department_id) VALUES 
('admin@kreatif.id', '$2a$10$7v2t5y.9V7m/3P9.1X.Y.O/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y', 'Super Admin', 1, NULL),
('manager@kreatif.id', '$2a$10$7v2t5y.9V7m/3P9.1X.Y.O/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y', 'IT Manager', 2, '550e8400-e29b-41d4-a716-446655440002'),
('controller@kreatif.id', '$2a$10$7v2t5y.9V7m/3P9.1X.Y.O/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y', 'Document Controller', 3, NULL),
('user@kreatif.id', '$2a$10$7v2t5y.9V7m/3P9.1X.Y.O/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y/Y', 'Staff IT', 4, '550e8400-e29b-41d4-a716-446655440002')
ON CONFLICT (email) DO NOTHING;
