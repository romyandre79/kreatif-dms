-- name: CreateDocument :one
INSERT INTO documents (
    title, description, file_name, file_path, file_size, mime_type, 
    company_id, branch_id, department_id, owner_id, status, batch_id,
    sensitivity, metadata, type_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) RETURNING *;

-- name: GetDocument :one
SELECT d.*, dt.name as type_name
FROM documents d
LEFT JOIN document_types dt ON d.type_id = dt.id
WHERE d.id = $1 LIMIT 1;

-- name: GetDocumentWithDetails :one
SELECT 
    d.*,
    dt.name as type_name,
    c.name as company_name,
    b.name as branch_name,
    dept.name as department_name,
    r.name as rack_name,
    bx.name as box_name,
    o.name as ordner_name,
    u.full_name as owner_name,
    aw.rejection_reason,
    aw.decision_note as rejection_notes
FROM documents d
LEFT JOIN document_types dt ON d.type_id = dt.id
LEFT JOIN companies c ON d.company_id = c.id
LEFT JOIN branches b ON d.branch_id = b.id
LEFT JOIN departments dept ON d.department_id = dept.id
LEFT JOIN racks r ON d.rack_id = r.id
LEFT JOIN boxes bx ON d.box_id = bx.id
LEFT JOIN ordners o ON d.ordner_id = o.id
LEFT JOIN users u ON d.owner_id = u.id
LEFT JOIN LATERAL (
    SELECT rejection_reason, decision_note 
    FROM approval_workflows 
    WHERE entity_id = d.id AND status = 'rejected' 
    ORDER BY created_at DESC LIMIT 1
) aw ON TRUE
WHERE d.id = $1 LIMIT 1;

-- name: ListDocumentsByDepartment :many
SELECT * FROM documents 
WHERE department_id = $1 
ORDER BY created_at DESC;

-- name: UpdateDocumentOCR :exec
UPDATE documents 
SET extracted_text = $2, 
    metadata = $3,
    is_ocr_processed = true, 
    status = 'active', 
    updated_at = NOW() 
WHERE id = $1;

-- name: UpdateDocumentMetadata :exec
UPDATE documents 
SET metadata = $2, updated_at = NOW() 
WHERE id = $1;

-- name: GetDocumentsByBatch :many
SELECT * FROM documents WHERE batch_id = $1;

-- name: CreateBatch :one
INSERT INTO processing_batches (user_id, total_files) 
VALUES ($1, $2) RETURNING *;

-- name: GetBatch :one
SELECT * FROM processing_batches WHERE id = $1 LIMIT 1;

-- name: UpdateBatchProgress :exec
UPDATE processing_batches 
SET processed_files = processed_files + 1, 
    status = CASE WHEN processed_files + 1 >= total_files THEN 'completed' ELSE 'processing' END,
    updated_at = NOW() 
WHERE id = $1;

-- name: CreateOCRJob :one
INSERT INTO ocr_jobs (
    entity_type, entity_id, ocr_service_url, ocr_engine, 
    source_file_path, raw_text, word_count, confidence_avg, words_json,
    status, processing_time_ms, preview_path, preview_paths, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW()
) RETURNING *;

-- name: GetOCRJobByEntity :one
SELECT * FROM ocr_jobs
WHERE entity_type = $1 AND entity_id = $2
ORDER BY created_at DESC
LIMIT 1;
-- name: ListOCRJobs :many
SELECT 
    j.id,
    j.entity_id,
    j.status,
    j.processing_time_ms,
    j.confidence_avg,
    j.preview_path,
    j.created_at,
    COALESCE(d.file_name, j.source_file_path) as filename,
    COALESCE(d.file_size, 0)::bigint as file_size,
    u.full_name as owner_name
FROM ocr_jobs j
LEFT JOIN documents d ON j.entity_id = d.id
LEFT JOIN users u ON d.owner_id = u.id
ORDER BY j.created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountOCRJobs :one
SELECT COUNT(*) FROM ocr_jobs;

-- name: ListRecentDocuments :many
SELECT 
    d.id, d.title, d.status, d.created_at, d.mime_type, d.file_size, d.metadata,
    dt.name as type_name,
    dept.name as department_name,
    COALESCE(d.description, '')::text as category
FROM documents d
LEFT JOIN document_types dt ON d.type_id = dt.id
LEFT JOIN departments dept ON d.department_id = dept.id
ORDER BY d.created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListRecentDocumentsByOwner :many
SELECT 
    d.id, d.title, d.status, d.created_at, d.mime_type, d.file_size, d.metadata,
    dt.name as type_name,
    dept.name as department_name,
    COALESCE(d.description, '')::text as category
FROM documents d
LEFT JOIN document_types dt ON d.type_id = dt.id
LEFT JOIN departments dept ON d.department_id = dept.id
WHERE d.owner_id = $1
ORDER BY d.created_at DESC
LIMIT $2 OFFSET $3;
-- name: GetDocumentLoanHistory :many
SELECT 
    br.id,
    u.full_name as user_name,
    u.department_id,
    dept.name as department_name,
    br.borrow_date,
    br.return_date,
    br.status,
    COALESCE(br.reason, '')::text as reason
FROM borrow_requests br
JOIN users u ON br.user_id = u.id
LEFT JOIN departments dept ON u.department_id = dept.id
WHERE br.document_id = $1
ORDER BY br.created_at DESC;
-- name: UpdateDocumentStatus :exec
UPDATE documents 
SET status = $2, updated_at = NOW() 
WHERE id = $1;

-- name: UpdateTaskStatusByEntity :exec
UPDATE approval_workflows 
SET status = $2,
    decided_at = NULL,
    decision_note = NULL,
    rejection_reason = NULL,
    updated_at = NOW()
WHERE entity_id = $1;

-- name: UpdateDocument :one
UPDATE documents 
SET title = $2, 
    description = $3,
    type_id = $4,
    sensitivity = $5,
    metadata = $6,
    file_name = COALESCE(NULLIF(@file_name::text, ''), file_name),
    file_path = COALESCE(NULLIF(@file_path::text, ''), file_path),
    file_size = CASE WHEN @file_size::bigint > 0 THEN @file_size::bigint ELSE file_size END,
    mime_type = COALESCE(NULLIF(@mime_type::text, ''), mime_type),
    status = COALESCE(NULLIF(@status::text, ''), 'pending'),
    updated_at = NOW() 
WHERE id = $1 
RETURNING *;
