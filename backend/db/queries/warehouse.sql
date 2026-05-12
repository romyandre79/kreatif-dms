-- name: ListDocumentsForLabeling :many
SELECT 
    d.id as document_id,
    d.title,
    dept.name as department_name,
    m.manifest_no,
    mi.status as manifest_item_status,
    d.physical_status,
    d.created_at
FROM documents d
JOIN physical_manifest_items mi ON d.id = mi.document_id
JOIN physical_manifests m ON mi.manifest_id = m.id
LEFT JOIN departments dept ON d.department_id = dept.id
WHERE mi.status = $1
ORDER BY d.created_at DESC;

-- name: UpdateManifestItemStatusBulk :exec
UPDATE physical_manifest_items
SET status = $2, verified_at = NOW()
WHERE document_id = ANY($1::uuid[]);

-- name: UpdateDocumentPhysicalStatusBulk :exec
UPDATE documents
SET physical_status = $2, updated_at = NOW()
WHERE id = ANY($1::uuid[]);

-- name: GetLabelingStats :one
SELECT 
    COUNT(*) FILTER (WHERE status = 'digitized') as waiting_count,
    COUNT(*) FILTER (WHERE status = 'labeled') as printed_count,
    COUNT(*) FILTER (WHERE status = 'archived') as stored_count
FROM physical_manifest_items;
