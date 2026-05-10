-- 000048_add_type_id_to_documents.up.sql

ALTER TABLE documents ADD COLUMN IF NOT EXISTS type_id UUID REFERENCES document_types(id) ON DELETE SET NULL;
