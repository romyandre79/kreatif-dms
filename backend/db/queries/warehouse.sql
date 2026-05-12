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
SET 
    physical_status = $2, 
    box_id = COALESCE(NULLIF(@box_id::text, '')::uuid, box_id),
    updated_at = NOW()
WHERE id = ANY($1::uuid[]);

-- name: GetLabelingStats :one
SELECT 
    COUNT(*) FILTER (WHERE status = 'digitized') as waiting_count,
    COUNT(*) FILTER (WHERE status = 'labeled') as printed_count,
    COUNT(*) FILTER (WHERE status = 'archived') as stored_count
FROM physical_manifest_items;

-- name: SearchBoxes :many
SELECT 
    bx.id, bx.name, r.name as rack_name, d.name as department_name,
    c.name as company_name, b.name as branch_name,
    r.location_detail
FROM boxes bx
JOIN racks r ON bx.rack_id = r.id
JOIN departments d ON r.department_id = d.id
JOIN branches b ON d.branch_id = b.id
JOIN companies c ON b.company_id = c.id
WHERE bx.name ILIKE '%' || $1 || '%' 
   OR r.name ILIKE '%' || $1 || '%'
ORDER BY bx.name
LIMIT 20;

-- name: FindEligibleBoxes :many
SELECT 
    bx.id, bx.name, r.id as rack_id, r.name as rack_name,
    d.id as department_id, d.name as department_name,
    bx.max_docs_capacity,
    (SELECT COUNT(*) FROM documents WHERE box_id = bx.id) as current_docs,
    r.location_detail
FROM boxes bx
JOIN racks r ON bx.rack_id = r.id
JOIN departments d ON r.department_id = d.id
WHERE d.id = @department_id
  AND (r.allowed_category_ids IS NULL OR cardinality(r.allowed_category_ids) = 0 OR @category_id::uuid = ANY(r.allowed_category_ids))
  AND (SELECT COUNT(*) FROM documents WHERE box_id = bx.id) < bx.max_docs_capacity
ORDER BY (bx.max_docs_capacity - (SELECT COUNT(*) FROM documents WHERE box_id = bx.id)) ASC
LIMIT 5;

