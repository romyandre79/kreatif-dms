-- name: GetDailyStats :one
SELECT 
    (SELECT COUNT(*) FROM documents WHERE created_at >= CURRENT_DATE) as daily_received,
    (SELECT COUNT(*) FROM ocr_jobs WHERE completed_at >= CURRENT_DATE AND status = 'completed') as daily_scanned,
    (SELECT COUNT(*) FROM documents WHERE updated_at >= CURRENT_DATE AND status = 'active') as daily_processed,
    (SELECT COALESCE(SUM(file_size), 0) FROM documents) as total_storage_size,
    (SELECT COUNT(*) FROM approval_workflows WHERE status = 'pending') as active_tasks;

-- name: GetUserDailyStats :one
SELECT 
    (SELECT COUNT(*) FROM documents WHERE owner_id = $1 AND created_at >= CURRENT_DATE) as daily_received,
    (SELECT COUNT(*) FROM borrow_requests WHERE user_id = $1 AND status = 'active') as active_loans,
    (SELECT COUNT(*) FROM approval_workflows WHERE approver_id = $1 AND status = 'pending') as active_tasks,
    (SELECT COUNT(*) FROM activity_logs WHERE user_id = $1 AND action = 'SEARCH' AND created_at >= CURRENT_DATE) as search_count;

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

-- name: GetUserRecentActivities :many
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
WHERE a.user_id = $1
ORDER BY a.created_at DESC
LIMIT $2;

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

-- name: GetUserPriorityTasks :many
SELECT 
    id,
    entity_type,
    entity_id,
    level,
    status,
    created_at
FROM approval_workflows
WHERE status = 'pending' AND (approver_id = $1)
ORDER BY created_at ASC
LIMIT $2;

-- name: GetUserLoanHistory :many
SELECT 
    br.id,
    d.title as document_title,
    br.borrow_date,
    br.due_date,
    br.status
FROM borrow_requests br
JOIN documents d ON br.document_id = d.id
WHERE br.user_id = $1
ORDER BY br.created_at DESC
LIMIT $2;
