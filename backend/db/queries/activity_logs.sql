-- name: CreateActivityLog :one
INSERT INTO activity_logs (
    user_id, action, entity_type, entity_id, details, ip_address
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: ListActivityLogs :many
SELECT 
    al.*, 
    u.full_name as user_name,
    u.email as user_email
FROM activity_logs al
JOIN users u ON al.user_id = u.id
ORDER BY al.created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetActivityLogsByEntity :many
SELECT 
    al.*, 
    u.full_name as user_name
FROM activity_logs al
JOIN users u ON al.user_id = u.id
WHERE al.entity_type = $1 AND al.entity_id = $2
ORDER BY al.created_at DESC;

-- name: GetZonationLogs :many
SELECT 
    al.id,
    al.created_at,
    al.action,
    al.details,
    u.full_name as user_name,
    COALESCE(r.name, '')::VARCHAR as rack_name,
    COALESCE(d.name, '')::VARCHAR as department_name
FROM activity_logs al
JOIN users u ON al.user_id = u.id
LEFT JOIN racks r ON al.entity_id = r.id
LEFT JOIN departments d ON (CASE WHEN al.details->>'department_id' IS NOT NULL THEN (al.details->>'department_id')::UUID ELSE NULL END) = d.id
WHERE al.entity_type = 'rack' AND al.action = 'UPDATE'
ORDER BY al.created_at DESC
LIMIT $1 OFFSET $2;
