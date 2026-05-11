-- 000055_add_physical_status_to_documents.down.sql
ALTER TABLE documents DROP COLUMN IF EXISTS physical_status;
