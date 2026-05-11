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
