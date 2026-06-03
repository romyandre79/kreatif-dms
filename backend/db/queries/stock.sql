-- name: GetStockMissionStats :one
SELECT 
    COUNT(CASE WHEN status = 'pending' THEN 1 END)::int as pending_missions,
    COUNT(CASE WHEN status = 'in_progress' THEN 1 END)::int as active_missions,
    COALESCE(
        ROUND(
            (COUNT(CASE WHEN status = 'completed' THEN 1 END)::float / 
             NULLIF(COUNT(*), 0) * 100)::numeric, 1
        )::float, 0.0
    )::float as accuracy_rate
FROM stock_opname_missions;

-- name: ListStockMissions :many
SELECT 
    m.id, m.session_id, m.title, m.description, m.assigned_to, m.target_area, m.status, m.completed_at, m.created_at,
    u.full_name as assigned_operator_name,
    COALESCE(s.session_no, '') as session_no,
    COALESCE(s.status, '') as session_status
FROM stock_opname_missions m
JOIN users u ON m.assigned_to = u.id
LEFT JOIN stock_opname_sessions s ON m.session_id = s.id
ORDER BY m.created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetStockMission :one
SELECT 
    m.id, m.session_id, m.title, m.description, m.assigned_to, m.target_area, m.status, m.completed_at, m.created_at,
    u.full_name as assigned_operator_name,
    s.session_no, s.status as session_status
FROM stock_opname_missions m
JOIN users u ON m.assigned_to = u.id
LEFT JOIN stock_opname_sessions s ON m.session_id = s.id
WHERE m.id = $1;

-- name: CreateStockMission :one
INSERT INTO stock_opname_missions (id, session_id, title, description, assigned_to, target_area, status, created_at)
VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, 'pending', NOW())
RETURNING *;

-- name: UpdateStockMissionStatus :exec
UPDATE stock_opname_missions 
SET status = $2, completed_at = CASE WHEN $2 = 'completed' THEN NOW() ELSE completed_at END
WHERE id = $1;

-- name: CreateStockSession :one
INSERT INTO stock_opname_sessions (id, session_no, branch_id, department_id, conducted_by, status, created_at, updated_at)
VALUES (gen_random_uuid(), $1, $2, $3, $4, 'in_progress', NOW(), NOW())
RETURNING *;

-- name: GetStockSession :one
SELECT * FROM stock_opname_sessions WHERE id = $1;

-- name: UpdateStockSessionSummary :exec
UPDATE stock_opname_sessions
SET total_matched = $2, total_on_loan = $3, total_missing = $4, total_extra = $5, updated_at = NOW()
WHERE id = $1;

-- name: UpdateStockSessionStatus :exec
UPDATE stock_opname_sessions
SET status = $2, approved_by = $3, approved_at = CASE WHEN $2 = 'approved' THEN NOW() ELSE approved_at END, resolution_note = $4, updated_at = NOW()
WHERE id = $1;

-- name: ListStockItemsBySession :many
SELECT * FROM stock_opname_items WHERE session_id = $1 ORDER BY scanned_at DESC;

-- name: GetStockItemBySku :one
SELECT * FROM stock_opname_items WHERE session_id = $1 AND sku_code = $2 LIMIT 1;

-- name: CreateStockItem :one
INSERT INTO stock_opname_items (id, session_id, document_id, sku_code, item_name, system_qty, physical_qty, status, location_name, sub_location, scanned_at)
VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
RETURNING *;

-- name: UpdateStockItemQuantity :exec
UPDATE stock_opname_items
SET physical_qty = $2, status = $3, scanned_at = NOW()
WHERE id = $1;

-- name: UpdateStockItemResolution :exec
UPDATE stock_opname_items
SET resolution = $2, resolution_note = $3
WHERE id = $1;

-- name: GetDocumentsByTargetArea :many
SELECT d.id, d.title, d.file_name, r.name as rack_name, b.name as box_name
FROM documents d
LEFT JOIN racks r ON d.rack_id = r.id
LEFT JOIN boxes b ON d.box_id = b.id
WHERE r.name ILIKE $1 OR b.name ILIKE $1 OR $1 = 'all';
