-- name: GetDailyStats :one
SELECT 
    (SELECT COUNT(*) FROM documents WHERE created_at >= CURRENT_DATE) as daily_received,
    (SELECT COUNT(*) FROM ocr_jobs WHERE completed_at >= CURRENT_DATE AND status = 'completed') as daily_scanned,
    (SELECT COUNT(*) FROM documents WHERE updated_at >= CURRENT_DATE AND status = 'active') as daily_processed,
    (SELECT COALESCE(SUM(file_size), 0) FROM documents) as total_storage_size,
    (SELECT COUNT(*) FROM approval_workflows WHERE status = 'pending') as active_tasks;

-- name: GetRecentActivities :many
SELECT 
    a.id,
    a.action,
    a.entity_type,
    a.entity_id,
    a.details,
    a.created_at,
    u.full_name as user_name,
    u.email as user_email
FROM activity_logs a
JOIN users u ON a.user_id = u.id
ORDER BY a.created_at DESC
LIMIT $1;

-- name: GetPriorityTasks :many
SELECT 
    id,
    entity_type,
    entity_id,
    level,
    status,
    created_at
FROM approval_workflows
WHERE status = 'pending'
ORDER BY created_at ASC
LIMIT $1;
