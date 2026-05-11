-- 000034_make_document_file_columns_nullable.down.sql
ALTER TABLE documents ALTER COLUMN file_name SET NOT NULL;
ALTER TABLE documents ALTER COLUMN file_path SET NOT NULL;
ALTER TABLE documents ALTER COLUMN file_size SET NOT NULL;
ALTER TABLE documents ALTER COLUMN mime_type SET NOT NULL;
