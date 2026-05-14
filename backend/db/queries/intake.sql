-- name: GetDocumentByShortID :one
SELECT d.*, u.full_name as owner_name, dept.name as department_name, dt.name as type_name
FROM documents d
JOIN users u ON d.owner_id = u.id
LEFT JOIN departments dept ON d.department_id = dept.id
LEFT JOIN document_types dt ON d.type_id = dt.id
WHERE d.id::text ILIKE $1 || '%'
LIMIT 1;

-- name: CreatePhysicalManifest :one
INSERT INTO physical_manifests (
    manifest_no, sender_id, department_id, total_items, status, notes
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: AddManifestItem :one
INSERT INTO physical_manifest_items (
    manifest_id, document_id, status, notes
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateManifestStatus :exec
UPDATE physical_manifests
SET status = $2, received_by = $3, received_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetIntakeStats :one
SELECT 
    COUNT(*) FILTER (WHERE physical_status = 'received' AND updated_at::date = CURRENT_DATE) as received_today,
    COUNT(*) FILTER (WHERE physical_status = 'pending') as pending_count,
    COUNT(*) FILTER (WHERE physical_status = 'rejected' AND updated_at::date = CURRENT_DATE) as rejected_today
FROM documents;


-- name: UpdateDocumentPhysicalStatus :exec
UPDATE documents
SET physical_status = $2, current_manifest_id = $3, updated_at = NOW()
WHERE id = $1;

-- name: GetManifestByNo :one
SELECT m.*, u.full_name as sender_name, dept.name as department_name
FROM physical_manifests m
JOIN users u ON m.sender_id = u.id
JOIN departments dept ON m.department_id = dept.id
WHERE m.manifest_no = $1;

-- name: GetManifestItems :many
SELECT 
    mi.*, 
    d.title as document_title, 
    d.type_id as document_type_id,
    dt.name as document_type, 
    dt.category_id as document_category_id,
    d.created_at as document_date,
    d.physical_status as current_physical_status,
    COALESCE(oj.preview_path, d.file_path) as preview_path,
    d.metadata as document_metadata,
    d.extracted_text as extracted_text
FROM physical_manifest_items mi
JOIN documents d ON mi.document_id = d.id
LEFT JOIN document_types dt ON d.type_id = dt.id
LEFT JOIN ocr_jobs oj ON oj.entity_id = d.id AND oj.entity_type = 'document'
WHERE mi.manifest_id = $1;

-- name: UpdateManifestItemStatus :exec
UPDATE physical_manifest_items
SET status = $2, verified_at = NOW(), notes = $3
WHERE id = $1;

-- name: ListPendingManifests :many
SELECT m.*, u.full_name as sender_name, dept.name as department_name
FROM physical_manifests m
JOIN users u ON m.sender_id = u.id
JOIN departments dept ON m.department_id = dept.id
WHERE m.status = 'pending'
ORDER BY m.created_at DESC;

-- name: ListPendingDocumentsWithoutManifest :many
SELECT d.*, u.full_name as owner_name, dept.name as department_name, dt.name as type_name
FROM documents d
JOIN users u ON d.owner_id = u.id
JOIN departments dept ON d.department_id = dept.id
LEFT JOIN document_types dt ON d.type_id = dt.id
WHERE d.physical_status = 'pending' AND d.current_manifest_id IS NULL
ORDER BY d.created_at DESC;
-- name: GetManifest :one
SELECT m.*, u.full_name as sender_name, dept.name as department_name
FROM physical_manifests m
JOIN users u ON m.sender_id = u.id
JOIN departments dept ON m.department_id = dept.id
WHERE m.id = $1;

-- name: ListStagingManifests :many
SELECT m.*, u.full_name as sender_name, dept.name as department_name
FROM physical_manifests m
JOIN users u ON m.sender_id = u.id
JOIN departments dept ON m.department_id = dept.id
WHERE m.status = 'received'
ORDER BY m.received_at ASC;

-- name: GetStagingStats :one
SELECT 
    COUNT(*) FILTER (WHERE status = 'received') as total_staging,
    COUNT(*) FILTER (WHERE status = 'received' AND received_at < NOW() - INTERVAL '4 hours') as overdue_count,
    COUNT(*) FILTER (WHERE (status = 'digitized' OR status = 'completed') AND updated_at::date = CURRENT_DATE) as completed_today,
    (SELECT COUNT(*) FROM users u JOIN roles r ON u.role_id = r.id WHERE r.name ILIKE '%doc controller%' AND u.status = 'active') as active_dc_count
FROM physical_manifests;

-- name: GetActiveDocControllers :many
SELECT u.full_name, u.avatar_url
FROM users u
JOIN roles r ON u.role_id = r.id
WHERE r.name ILIKE '%doc controller%' AND u.status = 'active'
LIMIT 5;

-- name: UpdateDocumentIndexing :exec
UPDATE documents
SET 
    title = COALESCE(NULLIF($2, ''), title),
    type_id = $3,
    metadata = $4,
    physical_status = $5,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateManifestItemStatusByDocID :exec
UPDATE physical_manifest_items
SET status = $3, verified_at = NOW()
WHERE manifest_id = $1 AND document_id = $2;

-- name: GetManifestProgress :one
SELECT 
    COUNT(*) as total_items,
    COUNT(*) FILTER (WHERE status = 'digitized') as digitized_count
FROM physical_manifest_items
WHERE manifest_id = $1;
