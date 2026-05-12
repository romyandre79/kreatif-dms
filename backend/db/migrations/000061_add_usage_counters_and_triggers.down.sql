-- 000061_add_usage_counters_and_triggers.down.sql

DROP TRIGGER IF EXISTS trg_box_capacity_sync ON boxes;
DROP TRIGGER IF EXISTS trg_document_capacity_sync ON documents;

DROP FUNCTION IF EXISTS fn_sync_box_capacity();
DROP FUNCTION IF EXISTS fn_sync_document_capacity();

ALTER TABLE boxes DROP COLUMN IF EXISTS current_docs_count;
ALTER TABLE racks DROP COLUMN IF EXISTS current_boxes_count;
ALTER TABLE departments DROP COLUMN IF EXISTS current_docs_count;
ALTER TABLE branches DROP COLUMN IF EXISTS current_docs_count;
ALTER TABLE companies DROP COLUMN IF EXISTS current_docs_count;
