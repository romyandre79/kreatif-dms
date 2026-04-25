-- ============================================================================
-- Migration 000010 DOWN: Revert all tables and columns added in 000010
-- ============================================================================

-- Drop indexes on existing tables first
DROP INDEX IF EXISTS idx_loan_dept;
DROP INDEX IF EXISTS idx_loan_branch;
DROP INDEX IF EXISTS idx_loan_company;
DROP INDEX IF EXISTS idx_circ_dept;
DROP INDEX IF EXISTS idx_circ_branch;
DROP INDEX IF EXISTS idx_circ_company;
DROP INDEX IF EXISTS idx_docs_es_indexed;
DROP INDEX IF EXISTS idx_docs_retention_exp;
DROP INDEX IF EXISTS idx_docs_sensitivity;

-- Revert columns on existing tables
ALTER TABLE processing_batches DROP COLUMN IF EXISTS department_id;
ALTER TABLE processing_batches DROP COLUMN IF EXISTS intake_session_id;
ALTER TABLE borrow_requests DROP COLUMN IF EXISTS rfid_tag_id;
ALTER TABLE borrow_requests DROP COLUMN IF EXISTS l2_approved_at;
ALTER TABLE borrow_requests DROP COLUMN IF EXISTS l2_approved_by;
ALTER TABLE documents DROP COLUMN IF EXISTS es_indexed;
ALTER TABLE documents DROP COLUMN IF EXISTS minio_bucket;
ALTER TABLE documents DROP COLUMN IF EXISTS circulation_id;
ALTER TABLE documents DROP COLUMN IF EXISTS sensitivity;
ALTER TABLE documents DROP COLUMN IF EXISTS retention_expiry_date;
ALTER TABLE documents DROP COLUMN IF EXISTS retention_years;

-- Drop new tables (reverse order of creation to respect FK)
DROP TABLE IF EXISTS dashboard_stats_cache;
DROP TABLE IF EXISTS intake_sessions;
DROP TABLE IF EXISTS ocr_jobs;
DROP TABLE IF EXISTS file_storage_objects;
DROP TABLE IF EXISTS search_index_status;
DROP TABLE IF EXISTS async_task_queue;
DROP TABLE IF EXISTS archival_number_sequences;
DROP TABLE IF EXISTS integration_monitors;
DROP TABLE IF EXISTS import_jobs;
DROP TABLE IF EXISTS backup_records;
DROP TABLE IF EXISTS compliance_remediations;
DROP TABLE IF EXISTS compliance_checks;
DROP TABLE IF EXISTS document_tracking_events;
DROP TABLE IF EXISTS rack_capacities;
DROP TABLE IF EXISTS print_jobs;
DROP TABLE IF EXISTS generated_labels;
DROP TABLE IF EXISTS loan_extensions;
DROP TABLE IF EXISTS approval_workflows;
DROP TABLE IF EXISTS staging_documents;
DROP TABLE IF EXISTS document_registrations;
DROP TABLE IF EXISTS sso_sync_logs;
DROP TABLE IF EXISTS system_settings;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS stock_opname_missions;
DROP TABLE IF EXISTS stock_opname_items;
DROP TABLE IF EXISTS stock_opname_sessions;
DROP TABLE IF EXISTS destruction_logs;
DROP TABLE IF EXISTS retention_items;
DROP TABLE IF EXISTS retention_batches;
DROP TABLE IF EXISTS retention_policies;
DROP TABLE IF EXISTS loan_cart_items;
DROP TABLE IF EXISTS loan_request_items;

-- Drop columns added to loan_requests/circulations before dropping the tables
ALTER TABLE loan_requests DROP COLUMN IF EXISTS department_id;
ALTER TABLE loan_requests DROP COLUMN IF EXISTS branch_id;
ALTER TABLE loan_requests DROP COLUMN IF EXISTS company_id;
ALTER TABLE document_circulations DROP COLUMN IF EXISTS department_id;
ALTER TABLE document_circulations DROP COLUMN IF EXISTS branch_id;
ALTER TABLE document_circulations DROP COLUMN IF EXISTS company_id;

DROP TABLE IF EXISTS loan_requests;
DROP TABLE IF EXISTS distribution_records;
DROP TABLE IF EXISTS digital_signatures;
DROP TABLE IF EXISTS routing_slip_recipients;
DROP TABLE IF EXISTS document_circulations;
