-- 000036_configure_manager_menu.down.sql
-- (Optional cleanup)
DELETE FROM role_permissions WHERE module_id IN ('cat_approval', 'sub_docs', 'sub_loans', 'sub_ext', 'cat_loans', 'loan_req', 'loan_my', 'loan_soft', 'loan_hist', 'cat_submissions', 'sub_new', 'sub_status', 'notifications_menu');
DELETE FROM system_modules WHERE id IN ('cat_approval', 'sub_docs', 'sub_loans', 'sub_ext', 'cat_loans', 'loan_req', 'loan_my', 'loan_soft', 'loan_hist', 'cat_submissions', 'sub_new', 'sub_status', 'notifications_menu');
