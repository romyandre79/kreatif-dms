-- ============================================================================
-- Migration 000010: Create all tables required by frontend views
-- Covers: Circulation, Loans, Retention, Stock, Notifications, Settings,
--         Registration, Approvals, Tracking, Warehouse Labels, Compliance
-- ============================================================================

-- ===========================================
-- 1. DOCUMENT CIRCULATION (Surat Masuk/Keluar)
-- Views: circulation/incoming, checkout, checkin, inbox, pickup, receive
-- ===========================================

-- Incoming/outgoing mail registration
CREATE TABLE IF NOT EXISTS document_circulations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ref_no VARCHAR(100) UNIQUE NOT NULL,          -- e.g. LGL.EXT.2026.0045
    direction VARCHAR(20) NOT NULL DEFAULT 'incoming', -- 'incoming', 'outgoing'
    sender VARCHAR(500),
    recipient_org VARCHAR(500),
    subject VARCHAR(500) NOT NULL,
    doc_title VARCHAR(500),
    classification VARCHAR(50) DEFAULT 'normal',  -- 'normal', 'confidential', 'secret'
    nature VARCHAR(50) DEFAULT 'normal',          -- 'urgent', 'normal', 'confidential'
    priority VARCHAR(20) DEFAULT 'normal',        -- 'low', 'normal', 'high', 'critical'
    received_date TIMESTAMPTZ,
    document_id UUID REFERENCES documents(id) ON DELETE SET NULL,
    
    -- OCR validation fields
    ocr_quality_pct INT DEFAULT 0,
    completeness_pct INT DEFAULT 0,
    is_duplicate BOOLEAN DEFAULT false,
    
    -- Workflow state
    status VARCHAR(50) NOT NULL DEFAULT 'draft',  -- 'draft','registered','routing','labelled','signed','distributed','archived'
    registered_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_circ_ref ON document_circulations(ref_no);
CREATE INDEX IF NOT EXISTS idx_circ_status ON document_circulations(status);
CREATE INDEX IF NOT EXISTS idx_circ_direction ON document_circulations(direction);

-- Routing slip recipients for circulation
CREATE TABLE IF NOT EXISTS routing_slip_recipients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    circulation_id UUID NOT NULL REFERENCES document_circulations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id),
    order_seq INT NOT NULL DEFAULT 1,
    role_label VARCHAR(255),            -- e.g. 'Head of Legal Department'
    status VARCHAR(50) DEFAULT 'pending', -- 'pending','viewed','signed','rejected'
    signed_at TIMESTAMPTZ,
    note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_rsr_circ ON routing_slip_recipients(circulation_id);

-- Digital signatures for circulation documents
CREATE TABLE IF NOT EXISTS digital_signatures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entity_type VARCHAR(50) NOT NULL,       -- 'circulation', 'retention', 'loan', 'stock_opname'
    entity_id UUID NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    signature_type VARCHAR(20) DEFAULT 'digital', -- 'digital', 'upload', 'typed'
    signature_data TEXT,                    -- base64 or file path
    approval_status VARCHAR(20) DEFAULT 'approved', -- 'approved','conditional','clarify','rejected'
    note TEXT,
    ip_address INET,
    signed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_digsig_entity ON digital_signatures(entity_type, entity_id);

-- Distribution records for circulated documents
CREATE TABLE IF NOT EXISTS distribution_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    circulation_id UUID NOT NULL REFERENCES document_circulations(id) ON DELETE CASCADE,
    recipient_user_id UUID NOT NULL REFERENCES users(id),
    channel VARCHAR(50) DEFAULT 'email',    -- 'email', 'whatsapp', 'system'
    status VARCHAR(50) DEFAULT 'pending',   -- 'pending','sent','delivered','read'
    sent_at TIMESTAMPTZ,
    read_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_dist_circ ON distribution_records(circulation_id);

-- ===========================================
-- 2. DOCUMENT LOANS (Peminjaman Dokumen)
-- Views: loans/fast-track, loans/my, loans/cart, loans/checkout, loans/tracking
-- ===========================================

CREATE TABLE IF NOT EXISTS loan_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    request_no VARCHAR(100) UNIQUE NOT NULL,   -- e.g. LOAN-2026-03-00125
    user_id UUID NOT NULL REFERENCES users(id),
    purpose VARCHAR(255) NOT NULL,              -- 'Internal Audit', 'External Audit', etc.
    department_filter VARCHAR(255),
    duration_days INT NOT NULL DEFAULT 3,
    notes TEXT,
    approver_notes TEXT,
    
    -- Approval workflow
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending','l1_approved','l2_approved','rejected','active','returned','overdue'
    l1_approved_by UUID REFERENCES users(id),
    l1_approved_at TIMESTAMPTZ,
    l1_rejection_reason TEXT,
    l2_approved_by UUID REFERENCES users(id),
    l2_approved_at TIMESTAMPTZ,
    l2_rejection_reason TEXT,
    
    -- Dates
    borrow_date TIMESTAMPTZ,
    due_date TIMESTAMPTZ,
    return_date TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_loan_user ON loan_requests(user_id);
CREATE INDEX IF NOT EXISTS idx_loan_status ON loan_requests(status);

-- Items within a loan request
CREATE TABLE IF NOT EXISTS loan_request_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    loan_request_id UUID NOT NULL REFERENCES loan_requests(id) ON DELETE CASCADE,
    document_id UUID NOT NULL REFERENCES documents(id),
    category VARCHAR(50),                   -- 'mandatory', 'recommended', 'sensitive'
    sensitivity VARCHAR(50),                -- 'CONFIDENTIAL', 'INTERNAL', 'RESTRICTED'
    reason TEXT,
    method VARCHAR(20) DEFAULT 'physical',  -- 'physical', 'digital'
    status VARCHAR(50) DEFAULT 'pending',   -- 'pending','picked_up','returned'
    picked_up_at TIMESTAMPTZ,
    returned_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_lri_loan ON loan_request_items(loan_request_id);

-- Loan cart (temporary selections before submission)
CREATE TABLE IF NOT EXISTS loan_cart_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    document_id UUID NOT NULL REFERENCES documents(id),
    category VARCHAR(50),
    added_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, document_id)
);

-- ===========================================
-- 3. RETENTION & DISPOSAL
-- Views: retention/approaching, batch, decision, export, history, 
--        log-detail, purge-confirm, purge-success, shredding, upload-bast
-- ===========================================

-- Retention policies per document category
CREATE TABLE IF NOT EXISTS retention_policies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    retention_years INT NOT NULL DEFAULT 5,
    department_id UUID REFERENCES departments(id),
    document_category VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Retention batches (group of docs approaching/past retention)
CREATE TABLE IF NOT EXISTS retention_batches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    batch_no VARCHAR(100) UNIQUE NOT NULL,  -- e.g. RET-2023-089
    department_id UUID REFERENCES departments(id),
    total_items INT NOT NULL DEFAULT 0,
    total_volume_m3 DECIMAL(10,2),
    risk_level VARCHAR(20) DEFAULT 'low',    -- 'low','medium','high'
    
    -- Decision
    decision VARCHAR(20),                    -- 'extended','destroy'
    decision_ref_no VARCHAR(255),            -- BAST reference number
    decision_by UUID REFERENCES users(id),
    decision_at TIMESTAMPTZ,
    decision_attachment TEXT,                 -- file path for BAST upload
    
    -- Process stages
    status VARCHAR(50) NOT NULL DEFAULT 'approaching', -- 'approaching','review','decision','shredding','completed'
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_retbatch_status ON retention_batches(status);

-- Individual items in a retention batch
CREATE TABLE IF NOT EXISTS retention_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    batch_id UUID NOT NULL REFERENCES retention_batches(id) ON DELETE CASCADE,
    document_id UUID NOT NULL REFERENCES documents(id),
    expiry_date DATE NOT NULL,
    status VARCHAR(50) DEFAULT 'active',     -- 'active','extended','destroyed'
    destroyed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_retitem_batch ON retention_items(batch_id);
CREATE INDEX IF NOT EXISTS idx_retitem_expiry ON retention_items(expiry_date);

-- Shredding/destruction log
CREATE TABLE IF NOT EXISTS destruction_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    batch_id UUID NOT NULL REFERENCES retention_batches(id),
    method VARCHAR(50) NOT NULL,             -- 'shredding', 'incineration', 'digital_wipe'
    witness_user_id UUID REFERENCES users(id),
    bast_file_path TEXT,
    notes TEXT,
    destroyed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ===========================================
-- 4. STOCK OPNAME / RECONCILIATION
-- Views: stock/scan, stock/reconciliation, stock/missions
-- ===========================================

CREATE TABLE IF NOT EXISTS stock_opname_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_no VARCHAR(100) UNIQUE NOT NULL,  -- e.g. #REC-2024-001
    branch_id UUID REFERENCES branches(id),
    department_id UUID REFERENCES departments(id),
    conducted_by UUID NOT NULL REFERENCES users(id),
    
    -- Summary counters
    total_matched INT DEFAULT 0,
    total_on_loan INT DEFAULT 0,
    total_missing INT DEFAULT 0,
    total_extra INT DEFAULT 0,
    
    status VARCHAR(50) NOT NULL DEFAULT 'in_progress', -- 'in_progress','pending_approval','approved','rejected'
    approved_by UUID REFERENCES users(id),
    approved_at TIMESTAMPTZ,
    resolution_note TEXT,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_so_status ON stock_opname_sessions(status);

-- Individual items scanned during stock opname
CREATE TABLE IF NOT EXISTS stock_opname_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID NOT NULL REFERENCES stock_opname_sessions(id) ON DELETE CASCADE,
    document_id UUID REFERENCES documents(id),
    sku_code VARCHAR(100),
    item_name VARCHAR(500),
    system_qty INT NOT NULL DEFAULT 0,
    physical_qty INT NOT NULL DEFAULT 0,
    variance INT GENERATED ALWAYS AS (physical_qty - system_qty) STORED,
    status VARCHAR(50) NOT NULL,              -- 'MATCH','MISSING','EXTRA','ON LOAN'
    location_name VARCHAR(255),
    sub_location VARCHAR(255),
    resolution VARCHAR(50),                   -- 'relocated','confirmed_lost','found','adjusted'
    resolution_note TEXT,
    scanned_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_soi_session ON stock_opname_items(session_id);

-- Stock opname missions/tasks
CREATE TABLE IF NOT EXISTS stock_opname_missions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID REFERENCES stock_opname_sessions(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    assigned_to UUID NOT NULL REFERENCES users(id),
    target_area VARCHAR(255),                 -- e.g. 'Rack A1-A5'
    status VARCHAR(50) DEFAULT 'pending',     -- 'pending','in_progress','completed'
    completed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ===========================================
-- 5. NOTIFICATIONS
-- Views: notifications/index
-- ===========================================

CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(500) NOT NULL,
    body TEXT,
    type VARCHAR(50) NOT NULL,                -- 'loan_approved','circulation','retention','system','approval_needed'
    entity_type VARCHAR(50),                  -- 'loan_request','circulation','document','retention_batch'
    entity_id UUID,
    channel VARCHAR(50) DEFAULT 'system',     -- 'system','email','whatsapp'
    is_read BOOLEAN DEFAULT false,
    read_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_notif_user ON notifications(user_id);
CREATE INDEX IF NOT EXISTS idx_notif_read ON notifications(user_id, is_read);

-- ===========================================
-- 6. SETTINGS (SSO / LDAP / System Config)
-- Views: settings/sso, settings/index
-- ===========================================

CREATE TABLE IF NOT EXISTS system_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category VARCHAR(100) NOT NULL,           -- 'sso', 'ldap', 'email', 'whatsapp', 'general'
    key VARCHAR(255) NOT NULL,
    value TEXT,
    value_type VARCHAR(20) DEFAULT 'string',  -- 'string','boolean','integer','json'
    description TEXT,
    updated_by UUID REFERENCES users(id),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(category, key)
);

-- LDAP/SSO sync logs
CREATE TABLE IF NOT EXISTS sso_sync_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    provider VARCHAR(50) NOT NULL,            -- 'active_directory', 'openldap'
    status VARCHAR(50) NOT NULL,              -- 'success', 'partial', 'failed'
    users_synced INT DEFAULT 0,
    groups_synced INT DEFAULT 0,
    errors INT DEFAULT 0,
    error_details JSONB DEFAULT '[]',
    started_at TIMESTAMPTZ NOT NULL,
    completed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ===========================================
-- 7. DOCUMENT REGISTRATION / MIGRATION
-- Views: registration/migration, registration/staging
-- ===========================================

-- Document registration queue (for legacy migration)
CREATE TABLE IF NOT EXISTS document_registrations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_filename VARCHAR(500) NOT NULL,
    source_pages INT,
    doc_type VARCHAR(100),                    -- Surat Perjanjian, Invoice, etc.
    doc_number VARCHAR(255),
    
    -- OCR extracted & verified data
    ocr_title VARCHAR(500),
    ocr_accuracy_pct INT DEFAULT 0,
    verified_title VARCHAR(500),
    
    -- Mandatory fields after validation
    department_id UUID REFERENCES departments(id),
    
    -- Duplicate check
    is_duplicate_checked BOOLEAN DEFAULT false,
    duplicate_document_id UUID REFERENCES documents(id),
    
    -- Location assignment
    assigned_rack_id UUID REFERENCES racks(id),
    assigned_box_id UUID REFERENCES boxes(id),
    assigned_ordner_id UUID REFERENCES ordners(id),
    capacity_score_pct INT,                   -- Location capacity score
    
    -- Final output
    final_filename VARCHAR(500),              -- e.g. DMS-INV-2024-00892-CORP.pdf
    final_document_id UUID REFERENCES documents(id),
    
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending','ocr_review','validated','location_assigned','registered','archived'
    registered_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_docreg_status ON document_registrations(status);

-- Staging area for bulk imports
CREATE TABLE IF NOT EXISTS staging_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    registration_id UUID REFERENCES document_registrations(id) ON DELETE SET NULL,
    original_filename VARCHAR(500) NOT NULL,
    file_path TEXT NOT NULL,
    file_size BIGINT,
    mime_type VARCHAR(100),
    status VARCHAR(50) DEFAULT 'pending',     -- 'pending','processing','completed','error'
    error_message TEXT,
    uploaded_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ===========================================
-- 8. APPROVAL WORKFLOWS (Generic)
-- Views: approvals/index, approvals/loans, approvals/extensions
-- ===========================================

CREATE TABLE IF NOT EXISTS approval_workflows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entity_type VARCHAR(50) NOT NULL,         -- 'loan_request','document_upload','retention','circulation'
    entity_id UUID NOT NULL,
    level INT NOT NULL DEFAULT 1,             -- L1, L2, etc.
    
    approver_id UUID NOT NULL REFERENCES users(id),
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending','approved','rejected','escalated'
    decision_note TEXT,
    rejection_reason TEXT,
    
    requires_pin BOOLEAN DEFAULT false,
    pin_verified BOOLEAN DEFAULT false,
    
    decided_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_aw_entity ON approval_workflows(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_aw_approver ON approval_workflows(approver_id, status);

-- Loan extension requests
CREATE TABLE IF NOT EXISTS loan_extensions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    loan_request_id UUID NOT NULL REFERENCES loan_requests(id) ON DELETE CASCADE,
    requested_by UUID NOT NULL REFERENCES users(id),
    extension_days INT NOT NULL,
    reason TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',     -- 'pending','approved','rejected'
    approved_by UUID REFERENCES users(id),
    decided_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ===========================================
-- 9. WAREHOUSE / PHYSICAL STORAGE MANAGEMENT
-- Views: warehouse/structure, warehouse/labels, 
--        admin/warehouse/departments, topology, monitoring, etc.
-- ===========================================

-- QR/Barcode labels generated for physical documents/storage
CREATE TABLE IF NOT EXISTS generated_labels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    label_type VARCHAR(50) NOT NULL,          -- 'document', 'rack', 'box', 'ordner'
    label_code VARCHAR(255) UNIQUE NOT NULL,  -- e.g. DOC-2023-0842, RACK-A1-04
    entity_type VARCHAR(50) NOT NULL,         -- 'document', 'rack', 'box', 'ordner', 'circulation'
    entity_id UUID NOT NULL,
    qr_data TEXT,                             -- encoded QR content
    encoding_level VARCHAR(50) DEFAULT 'full', -- 'ref_only','full','url'
    
    -- Print info
    printer_name VARCHAR(255),
    copies INT DEFAULT 1,
    last_printed_at TIMESTAMPTZ,
    print_count INT DEFAULT 0,
    
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_label_code ON generated_labels(label_code);
CREATE INDEX IF NOT EXISTS idx_label_entity ON generated_labels(entity_type, entity_id);

-- Print job history
CREATE TABLE IF NOT EXISTS print_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    label_id UUID REFERENCES generated_labels(id),
    printer_name VARCHAR(255) NOT NULL,
    copies INT DEFAULT 1,
    status VARCHAR(50) NOT NULL DEFAULT 'queued', -- 'queued','printing','success','failed'
    error_message TEXT,
    printed_by UUID REFERENCES users(id),
    printed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Rack capacity & override tracking
CREATE TABLE IF NOT EXISTS rack_capacities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    rack_id UUID NOT NULL REFERENCES racks(id) ON DELETE CASCADE,
    max_capacity INT NOT NULL DEFAULT 100,
    current_usage INT NOT NULL DEFAULT 0,
    usage_pct DECIMAL(5,2) GENERATED ALWAYS AS (
        CASE WHEN max_capacity > 0 THEN (current_usage::DECIMAL / max_capacity * 100) ELSE 0 END
    ) STORED,
    last_audit_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(rack_id)
);

-- ===========================================
-- 10. TRACKING & AUDIT TRAIL
-- Views: tracking/index
-- ===========================================

-- Document tracking events (checkin/checkout/movement)
CREATE TABLE IF NOT EXISTS document_tracking_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    document_id UUID NOT NULL REFERENCES documents(id),
    event_type VARCHAR(50) NOT NULL,          -- 'checkout','checkin','transfer','scan','view','download'
    from_location VARCHAR(255),
    to_location VARCHAR(255),
    performed_by UUID NOT NULL REFERENCES users(id),
    rfid_tag_id VARCHAR(255) REFERENCES rfid_tags(tag_id),
    note TEXT,
    ip_address INET,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_dte_doc ON document_tracking_events(document_id);
CREATE INDEX IF NOT EXISTS idx_dte_type ON document_tracking_events(event_type);

-- ===========================================
-- 11. COMPLIANCE & SECURITY AUDITS
-- Views: admin/compliance
-- ===========================================

CREATE TABLE IF NOT EXISTS compliance_checks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    check_type VARCHAR(100) NOT NULL,         -- 'tls','ssl_cert','aes_encryption','patch','audit_sync'
    source VARCHAR(255),                      -- e.g. 'Node-04 Firewall'
    check_id VARCHAR(100),                    -- e.g. 'F57-002-C'
    result VARCHAR(50) NOT NULL,              -- 'PASSED','WARNING','CRITICAL','FAILED'
    details JSONB DEFAULT '{}',
    health_score_pct INT,
    checked_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_cc_type ON compliance_checks(check_type);

-- Compliance remediation tasks
CREATE TABLE IF NOT EXISTS compliance_remediations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    compliance_check_id UUID REFERENCES compliance_checks(id),
    severity VARCHAR(20) NOT NULL,            -- 'urgent','patch','hardening'
    title VARCHAR(500) NOT NULL,
    description TEXT,
    action_type VARCHAR(50),                  -- 'renew_cert','apply_patch','auto_fix'
    status VARCHAR(50) DEFAULT 'open',        -- 'open','in_progress','resolved','dismissed'
    resolved_by UUID REFERENCES users(id),
    resolved_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ===========================================
-- 12. ADMIN: BACKUP, IMPORT, MONITORING
-- Views: admin/backup, admin/import/excel, admin/monitoring/integration,
--        admin/deployment/readiness, admin/archival/numbering
-- ===========================================

-- Backup records
CREATE TABLE IF NOT EXISTS backup_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    backup_type VARCHAR(50) NOT NULL,         -- 'full','incremental','differential'
    target VARCHAR(100) NOT NULL,             -- 'database','files','minio','all'
    file_path TEXT,
    file_size BIGINT,
    status VARCHAR(50) NOT NULL DEFAULT 'running', -- 'running','completed','failed'
    error_message TEXT,
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ,
    created_by UUID REFERENCES users(id)
);

-- Excel/bulk import jobs
CREATE TABLE IF NOT EXISTS import_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_type VARCHAR(50) NOT NULL,         -- 'excel','csv','legacy_system'
    source_filename VARCHAR(500),
    total_rows INT DEFAULT 0,
    processed_rows INT DEFAULT 0,
    success_rows INT DEFAULT 0,
    error_rows INT DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending','processing','completed','failed'
    error_log JSONB DEFAULT '[]',
    uploaded_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

-- Integration monitoring (external system connections)
CREATE TABLE IF NOT EXISTS integration_monitors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name VARCHAR(255) NOT NULL,       -- 'elasticsearch','minio','redis','ocr_service','ldap'
    endpoint VARCHAR(500),
    status VARCHAR(50) NOT NULL DEFAULT 'unknown', -- 'healthy','degraded','down','unknown'
    response_time_ms INT,
    last_check_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    error_message TEXT,
    metadata JSONB DEFAULT '{}'
);

-- Archival numbering sequences
CREATE TABLE IF NOT EXISTS archival_number_sequences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prefix VARCHAR(50) NOT NULL,              -- e.g. 'LGL', 'FIN', 'HR'
    department_id UUID REFERENCES departments(id),
    current_seq INT NOT NULL DEFAULT 0,
    year INT NOT NULL,
    format_template VARCHAR(255) NOT NULL DEFAULT '{PREFIX}.{TYPE}.{YEAR}.{SEQ:4}',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(prefix, department_id, year)
);

-- ===========================================
-- 13. DOCUMENT INTAKE PIPELINE
-- Views: intake/duplicate-check, external-registration, final-review,
--        management, ocr-processing, ocr-review, path-review, scan-web
-- ===========================================

CREATE TABLE IF NOT EXISTS intake_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source VARCHAR(50) NOT NULL,              -- 'scanner','upload','email','web_scan','external'
    source_detail TEXT,                       -- scanner name, email address, URL, etc.
    
    -- Document info
    original_filename VARCHAR(500),
    file_path TEXT,
    file_size BIGINT,
    mime_type VARCHAR(100),
    page_count INT,
    
    -- External registration fields
    external_sender VARCHAR(500),
    external_ref_no VARCHAR(255),
    external_date DATE,
    
    -- OCR processing
    ocr_status VARCHAR(50) DEFAULT 'pending', -- 'pending','processing','completed','failed'
    ocr_text TEXT,
    ocr_confidence_pct INT,
    ai_refined_text TEXT,
    
    -- Review
    reviewed_by UUID REFERENCES users(id),
    review_status VARCHAR(50) DEFAULT 'pending', -- 'pending','approved','needs_correction'
    review_note TEXT,
    
    -- Duplicate check
    duplicate_status VARCHAR(50) DEFAULT 'unchecked', -- 'unchecked','unique','potential_duplicate','confirmed_duplicate'
    duplicate_document_id UUID REFERENCES documents(id),
    
    -- Path/location review
    suggested_department_id UUID REFERENCES departments(id),
    suggested_rack_id UUID REFERENCES racks(id),
    path_override BOOLEAN DEFAULT false,
    
    -- Final output
    final_document_id UUID REFERENCES documents(id),
    status VARCHAR(50) NOT NULL DEFAULT 'scanning', -- 'scanning','ocr','review','path_review','final_review','completed','rejected'
    
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_intake_status ON intake_sessions(status);

-- ===========================================
-- 14. DASHBOARD STATISTICS (Materialized/cache)
-- Views: dashboard/manager/review
-- ===========================================

CREATE TABLE IF NOT EXISTS dashboard_stats_cache (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stat_type VARCHAR(100) NOT NULL,          -- 'total_docs','pending_approvals','active_loans', etc.
    scope_type VARCHAR(50),                   -- 'global','branch','department'
    scope_id UUID,                            -- branch_id or department_id
    value DECIMAL(15,2) NOT NULL DEFAULT 0,
    metadata JSONB DEFAULT '{}',
    computed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(stat_type, scope_type, scope_id)
);

-- ===========================================
-- 15. SERVICE INTEGRATIONS
-- Relations with: Asynq (Redis worker), Elasticsearch, MinIO, OCR Service, LDAP
-- ===========================================

-- Async task queue tracking (relates to Asynq/Redis worker service)
-- Links every background job to the entity it processes
CREATE TABLE IF NOT EXISTS async_task_queue (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    task_type VARCHAR(100) NOT NULL,          -- 'document:ocr','intake:process','notification:send','backup:run','sso:sync','compliance:scan'
    asynq_task_id VARCHAR(255),              -- Asynq task ID from Redis
    entity_type VARCHAR(50) NOT NULL,         -- 'document','intake_session','loan_request','retention_batch','backup_record','import_job'
    entity_id UUID NOT NULL,                  -- FK resolved at app level (polymorphic)
    
    -- Worker execution info
    queue_name VARCHAR(50) DEFAULT 'default', -- 'default','critical','low'
    priority INT DEFAULT 0,
    max_retries INT DEFAULT 3,
    retry_count INT DEFAULT 0,
    
    status VARCHAR(50) NOT NULL DEFAULT 'queued', -- 'queued','processing','completed','failed','dead'
    error_message TEXT,
    
    -- Timing
    enqueued_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    
    -- Traceability
    enqueued_by UUID REFERENCES users(id),
    processed_by_worker VARCHAR(255)          -- worker hostname/ID
);

CREATE INDEX IF NOT EXISTS idx_atq_entity ON async_task_queue(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_atq_status ON async_task_queue(status);
CREATE INDEX IF NOT EXISTS idx_atq_type ON async_task_queue(task_type);

-- Elasticsearch index status tracking (relates to SearchService/Elasticsearch)
-- Tracks which documents have been indexed and their sync state
CREATE TABLE IF NOT EXISTS search_index_status (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    document_id UUID NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
    index_name VARCHAR(100) NOT NULL DEFAULT 'documents', -- ES index name
    es_doc_id VARCHAR(255),                 -- Elasticsearch document _id
    
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending','indexed','failed','stale','deleted'
    version INT DEFAULT 1,                  -- incremented on re-index
    
    indexed_at TIMESTAMPTZ,
    last_synced_at TIMESTAMPTZ,
    error_message TEXT,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(document_id, index_name)
);

CREATE INDEX IF NOT EXISTS idx_sis_status ON search_index_status(status);

-- MinIO file storage objects (relates to StorageService/MinIO)
-- Master registry of all files stored in MinIO with encryption status
CREATE TABLE IF NOT EXISTS file_storage_objects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bucket_name VARCHAR(100) NOT NULL DEFAULT 'dms-documents',
    object_key TEXT NOT NULL,               -- full path in MinIO: {dept_id}/{filename}
    
    -- Linked entity (polymorphic)
    entity_type VARCHAR(50) NOT NULL,        -- 'document','document_version','staging_document','intake_session','destruction_log','backup_record','digital_signature'
    entity_id UUID NOT NULL,
    
    -- File metadata
    file_name VARCHAR(500) NOT NULL,
    file_size BIGINT NOT NULL,
    mime_type VARCHAR(100),
    checksum VARCHAR(128),                  -- SHA256 hash
    
    -- Encryption status (AES-256 as per PROJECT_PLAN)
    is_encrypted BOOLEAN DEFAULT true,
    encryption_algorithm VARCHAR(50) DEFAULT 'AES-256-SSE',
    
    -- Lifecycle
    is_deleted BOOLEAN DEFAULT false,        -- soft delete
    deleted_at TIMESTAMPTZ,
    
    uploaded_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_fso_entity ON file_storage_objects(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_fso_bucket ON file_storage_objects(bucket_name, object_key);

-- OCR processing jobs (relates to OCR Service / PaddleOCR microservice)
-- Tracks every OCR invocation with results linked back to source entities
CREATE TABLE IF NOT EXISTS ocr_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Source entity (document, intake_session, or document_registration)
    entity_type VARCHAR(50) NOT NULL,        -- 'document','intake_session','document_registration'
    entity_id UUID NOT NULL,
    
    -- OCR service info
    ocr_service_url VARCHAR(500),            -- e.g. 'http://ocr-service:8000'
    ocr_engine VARCHAR(50) DEFAULT 'paddleocr', -- 'paddleocr','tesseract'
    
    -- Input
    source_file_path TEXT,                   -- MinIO object key
    source_pages INT,
    
    -- Results
    raw_text TEXT,                           -- raw OCR output
    word_count INT,
    confidence_avg DECIMAL(5,2),             -- average confidence %
    words_json JSONB,                        -- detailed word-level results [{text, confidence, page, box}]
    
    -- AI Refinement (relates to Gemini AI service)
    ai_provider VARCHAR(50),                 -- 'gemini','openai',null
    ai_refined_text TEXT,
    ai_metadata JSONB DEFAULT '{}',          -- extracted metadata from AI
    ai_refinement_status VARCHAR(50),        -- 'skipped','success','failed'
    
    -- Processing
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending','processing','completed','failed'
    error_message TEXT,
    processing_time_ms INT,
    
    -- Async task link
    async_task_id UUID REFERENCES async_task_queue(id),
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_ocr_entity ON ocr_jobs(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_ocr_status ON ocr_jobs(status);

-- ===========================================
-- 16. ADDITIONAL COLUMNS ON EXISTING TABLES
-- Strengthening org hierarchy relations and service links
-- ===========================================

-- Documents: add retention, sensitivity, circulation, and storage link fields
ALTER TABLE documents ADD COLUMN IF NOT EXISTS retention_years INT;
ALTER TABLE documents ADD COLUMN IF NOT EXISTS retention_expiry_date DATE;
ALTER TABLE documents ADD COLUMN IF NOT EXISTS sensitivity VARCHAR(50) DEFAULT 'internal'; -- 'public','internal','confidential','secret'
ALTER TABLE documents ADD COLUMN IF NOT EXISTS circulation_id UUID REFERENCES document_circulations(id) ON DELETE SET NULL;
ALTER TABLE documents ADD COLUMN IF NOT EXISTS minio_bucket VARCHAR(100) DEFAULT 'dms-documents';
ALTER TABLE documents ADD COLUMN IF NOT EXISTS es_indexed BOOLEAN DEFAULT false;

-- Document Circulations: tie to org hierarchy
ALTER TABLE document_circulations ADD COLUMN IF NOT EXISTS company_id UUID REFERENCES companies(id);
ALTER TABLE document_circulations ADD COLUMN IF NOT EXISTS branch_id UUID REFERENCES branches(id);
ALTER TABLE document_circulations ADD COLUMN IF NOT EXISTS department_id UUID REFERENCES departments(id);

-- Loan Requests: tie to org hierarchy
ALTER TABLE loan_requests ADD COLUMN IF NOT EXISTS company_id UUID REFERENCES companies(id);
ALTER TABLE loan_requests ADD COLUMN IF NOT EXISTS branch_id UUID REFERENCES branches(id);
ALTER TABLE loan_requests ADD COLUMN IF NOT EXISTS department_id UUID REFERENCES departments(id);

-- Borrow Requests (existing table): add L2 approval and RFID relation
ALTER TABLE borrow_requests ADD COLUMN IF NOT EXISTS l2_approved_by UUID REFERENCES users(id);
ALTER TABLE borrow_requests ADD COLUMN IF NOT EXISTS l2_approved_at TIMESTAMPTZ;
ALTER TABLE borrow_requests ADD COLUMN IF NOT EXISTS rfid_tag_id VARCHAR(255) REFERENCES rfid_tags(tag_id);

-- Processing Batches (existing): link to intake pipeline
ALTER TABLE processing_batches ADD COLUMN IF NOT EXISTS intake_session_id UUID;
ALTER TABLE processing_batches ADD COLUMN IF NOT EXISTS department_id UUID REFERENCES departments(id);

-- ===========================================
-- FINAL: Create indexes for new columns
-- ===========================================

CREATE INDEX IF NOT EXISTS idx_docs_sensitivity ON documents(sensitivity);
CREATE INDEX IF NOT EXISTS idx_docs_retention_exp ON documents(retention_expiry_date);
CREATE INDEX IF NOT EXISTS idx_docs_es_indexed ON documents(es_indexed);
CREATE INDEX IF NOT EXISTS idx_circ_company ON document_circulations(company_id);
CREATE INDEX IF NOT EXISTS idx_circ_branch ON document_circulations(branch_id);
CREATE INDEX IF NOT EXISTS idx_circ_dept ON document_circulations(department_id);
CREATE INDEX IF NOT EXISTS idx_loan_company ON loan_requests(company_id);
CREATE INDEX IF NOT EXISTS idx_loan_branch ON loan_requests(branch_id);
CREATE INDEX IF NOT EXISTS idx_loan_dept ON loan_requests(department_id);
