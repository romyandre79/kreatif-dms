-- name: CreateDocument :one
INSERT INTO documents (
    title, description, file_name, file_path, file_size, mime_type, 
    company_id, branch_id, department_id, owner_id, status, batch_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetDocument :one
SELECT * FROM documents WHERE id = $1 LIMIT 1;

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
    status, processing_time_ms, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW()
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
    id, title, status, created_at, mime_type, file_size, metadata,
    COALESCE(description, '')::text as category
FROM documents
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
