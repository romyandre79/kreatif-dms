-- 000036_configure_manager_menu.up.sql

-- 1. Create or Update modules for Manager role
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id) VALUES
('cat_approval', 'Approval Management', 'Manager', NULL, 'LucideCheckCircle2', ARRAY['VIEW'], 10, NULL),
('sub_docs', 'Document Submissions', 'Manager', '/approvals/submissions', 'LucideFileText', ARRAY['VIEW', 'APPROVE'], 11, 'cat_approval'),
('sub_loans', 'Document Loans', 'Manager', '/approvals/loans', 'LucideBookOpen', ARRAY['VIEW', 'APPROVE'], 12, 'cat_approval'),
('sub_ext', 'Loan Extensions', 'Manager', '/approvals/extensions', 'LucideHistory', ARRAY['VIEW', 'APPROVE'], 13, 'cat_approval'),

('search', 'Document Search', 'Search', '/documents', 'LucideSearch', ARRAY['VIEW'], 20, NULL),

('cat_loans', 'Loans', 'Circulation', NULL, 'LucideBookOpen', ARRAY['VIEW'], 30, NULL),
('loan_req', 'Request Loan', 'Circulation', '/loans/request', 'LucidePlusCircle', ARRAY['VIEW', 'CREATE'], 31, 'cat_loans'),
('loan_my', 'My Loans', 'Circulation', '/loans/my', 'LucideFileText', ARRAY['VIEW'], 32, 'cat_loans'),
('loan_soft', 'Softcopy Request', 'Circulation', '/loans/softcopy', 'LucideFileJson', ARRAY['VIEW', 'CREATE'], 33, 'cat_loans'),
('loan_hist', 'Loan History', 'Circulation', '/loans/history', 'LucideHistory', ARRAY['VIEW'], 34, 'cat_loans'),

('cat_submissions', 'Document Submissions', 'Registration', NULL, 'LucideLayers', ARRAY['VIEW'], 40, NULL),
('sub_new', 'Submit Document', 'Registration', '/registration/new', 'LucideUpload', ARRAY['VIEW', 'CREATE'], 41, 'cat_submissions'),
('sub_status', 'Submission Status', 'Registration', '/registration/status', 'LucideActivity', ARRAY['VIEW'], 42, 'cat_submissions'),

('notifications', 'Notifications', 'System', '/notifications', 'LucideBell', ARRAY['VIEW'], 100, NULL)
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    icon = EXCLUDED.icon,
    path = EXCLUDED.path,
    sort_order = EXCLUDED.sort_order,
    parent_id = EXCLUDED.parent_id;

-- 2. Configure Role Permissions for 'manajer'
DO $$
DECLARE
    manager_role_id INT;
BEGIN
    SELECT id INTO manager_role_id FROM roles WHERE name = 'manajer';
    
    IF manager_role_id IS NOT NULL THEN
        -- Clear existing permissions
        DELETE FROM role_permissions WHERE role_id = manager_role_id;
        
        -- Grant VIEW permission for the requested modules
        INSERT INTO role_permissions (role_id, module_id, action) VALUES
        (manager_role_id, 'dashboard', 'VIEW'),
        (manager_role_id, 'cat_approval', 'VIEW'),
        (manager_role_id, 'sub_docs', 'VIEW'),
        (manager_role_id, 'sub_loans', 'VIEW'),
        (manager_role_id, 'sub_ext', 'VIEW'),
        (manager_role_id, 'search', 'VIEW'),
        (manager_role_id, 'cat_loans', 'VIEW'),
        (manager_role_id, 'loan_req', 'VIEW'),
        (manager_role_id, 'loan_my', 'VIEW'),
        (manager_role_id, 'loan_soft', 'VIEW'),
        (manager_role_id, 'loan_hist', 'VIEW'),
        (manager_role_id, 'cat_submissions', 'VIEW'),
        (manager_role_id, 'sub_new', 'VIEW'),
        (manager_role_id, 'sub_status', 'VIEW'),
        (manager_role_id, 'notifications', 'VIEW');
    END IF;
END $$;
