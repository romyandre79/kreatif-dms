-- Insert default admin user if not exists
INSERT INTO users (email, password_hash, full_name, role_id, status)
SELECT 'admin@kreatif.id', '$2a$10$5/dIBIxr5ri3wR7xEy3w/.sLo9R09aO1hlCnFAu7x1Qx50e1uzGJC', 'Administrator', id, 'approved'
FROM roles WHERE name = 'admin'
ON CONFLICT (email) DO NOTHING;
