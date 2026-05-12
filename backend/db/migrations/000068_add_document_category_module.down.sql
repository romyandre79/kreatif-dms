-- Migration: Remove Document Category module
DELETE FROM role_permissions WHERE module_id = 'doc_category';
DELETE FROM system_modules WHERE id = 'doc_category';
