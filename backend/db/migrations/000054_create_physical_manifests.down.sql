-- 000054_create_physical_manifests.down.sql

ALTER TABLE documents DROP COLUMN IF EXISTS current_manifest_id;
DROP TABLE IF EXISTS physical_manifest_items;
DROP TABLE IF EXISTS physical_manifests;
