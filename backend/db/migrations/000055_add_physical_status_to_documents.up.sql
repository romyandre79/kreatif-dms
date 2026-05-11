-- 000055_add_physical_status_to_documents.up.sql
ALTER TABLE documents ADD COLUMN IF NOT EXISTS physical_status VARCHAR(50) DEFAULT 'none'; -- 'none', 'pending', 'received', 'archived', 'lost'

-- Update existing active documents to 'received' if they were already in the system
UPDATE documents SET physical_status = 'received' WHERE status = 'active';
-- Update pending documents (approved but not yet received)
UPDATE documents SET physical_status = 'pending' WHERE status = 'approved_pending_physical';
