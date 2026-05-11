-- 000034_make_document_file_columns_nullable.up.sql
ALTER TABLE documents ALTER COLUMN file_name DROP NOT NULL;
ALTER TABLE documents ALTER COLUMN file_path DROP NOT NULL;
ALTER TABLE documents ALTER COLUMN file_size DROP NOT NULL;
ALTER TABLE documents ALTER COLUMN mime_type DROP NOT NULL;
