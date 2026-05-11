-- name: GetDailyStats :one
SELECT 
    (SELECT COUNT(*) FROM documents WHERE created_at >= CURRENT_DATE) as daily_received,
    (SELECT COUNT(*) FROM ocr_jobs WHERE completed_at >= CURRENT_DATE AND status = 'completed') as daily_scanned,
    (SELECT COUNT(*) FROM documents WHERE updated_at >= CURRENT_DATE AND status = 'active') as daily_processed,
    (SELECT COALESCE(SUM(file_size), 0) FROM documents) as total_storage_size,
    (SELECT COUNT(*) FROM approval_workflows WHERE status = 'pending') as active_tasks,
    (SELECT COALESCE(AVG(EXTRACT(EPOCH FROM (decided_at - created_at))), 0)::float FROM approval_workflows WHERE status IN ('approved', 'rejected')) as avg_approval_time,
    (SELECT COALESCE((COUNT(*) FILTER (WHERE decided_at - created_at <= (SELECT COALESCE(NULLIF(value, ''), '24') FROM system_settings WHERE key = 'sla_approval_hours')::int * interval '1 hour')::float / NULLIF(COUNT(*), 0)::float) * 100, 100)::float FROM approval_workflows WHERE status IN ('approved', 'rejected')) as compliance_rate;

-- name: GetUserDailyStats :one
SELECT 
    (SELECT COUNT(*) FROM documents d1 WHERE d1.owner_id = $1 AND (d1.created_at AT TIME ZONE 'Asia/Jakarta')::date = (now() AT TIME ZONE 'Asia/Jakarta')::date) as daily_received,
    (SELECT COUNT(*) FROM borrow_requests br WHERE br.user_id = $1 AND br.status = 'active') as active_loans,
    (SELECT COUNT(*) FROM approval_workflows aw JOIN documents d2 ON aw.entity_id = d2.id WHERE d2.owner_id = $1 AND aw.status = 'pending') as active_tasks,
    (SELECT COUNT(*) FROM activity_logs al WHERE al.user_id = $1 AND al.action = 'SEARCH' AND (al.created_at AT TIME ZONE 'Asia/Jakarta')::date = (now() AT TIME ZONE 'Asia/Jakarta')::date) as search_count;

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
    aw.id,
    aw.entity_type,
    aw.entity_id,
    aw.level,
    aw.status,
    aw.created_at,
    COALESCE(d.title, '')::text as title,
    u.full_name as staff_name
FROM approval_workflows aw
LEFT JOIN documents d ON aw.entity_id = d.id
LEFT JOIN users u ON d.owner_id = u.id
WHERE aw.status = 'pending'
ORDER BY aw.created_at ASC
LIMIT $1;

-- name: GetUserPriorityTasks :many
SELECT 
    aw.id,
    aw.entity_type,
    aw.entity_id,
    aw.level,
    aw.status,
    aw.created_at,
    u.full_name as staff_name,
    COALESCE(d.title, '')::text as title
FROM approval_workflows aw
LEFT JOIN documents d ON aw.entity_id = d.id
LEFT JOIN users u ON d.owner_id = u.id
WHERE aw.status = 'pending' AND (aw.approver_id = $1)
UNION ALL
SELECT
    d.id,
    CASE WHEN d.status = 'draft' THEN 'document_draft' ELSE 'document_rejection' END as entity_type,
    d.id as entity_id,
    1 as level,
    d.status,
    d.created_at,
    u.full_name as staff_name,
    d.title
FROM documents d
LEFT JOIN users u ON d.owner_id = u.id
WHERE d.owner_id = $1 AND d.status IN ('rejected', 'draft')
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

-- name: GetManagerDailyStats :one
SELECT 
    (SELECT COUNT(*) FROM documents d 
     WHERE d.department_id IN (SELECT id FROM departments WHERE head_id = $1) 
     AND (d.created_at AT TIME ZONE 'Asia/Jakarta')::date = (now() AT TIME ZONE 'Asia/Jakarta')::date) as daily_received,
    (SELECT COUNT(*) FROM approval_workflows aw1 WHERE aw1.approver_id = $1 AND aw1.status = 'pending') as active_tasks,
    (SELECT COUNT(*) FROM approval_workflows aw2 WHERE aw2.approver_id = $1 AND aw2.status IN ('approved', 'rejected') 
     AND (aw2.decided_at AT TIME ZONE 'Asia/Jakarta')::date = (now() AT TIME ZONE 'Asia/Jakarta')::date) as daily_processed,
    (SELECT COUNT(*) FROM documents d WHERE d.owner_id = $1 AND (d.created_at AT TIME ZONE 'Asia/Jakarta')::date = (now() AT TIME ZONE 'Asia/Jakarta')::date) as my_submissions,
    (SELECT COALESCE(AVG(EXTRACT(EPOCH FROM (decided_at - created_at))), 0)::float FROM approval_workflows WHERE approver_id = $1 AND status IN ('approved', 'rejected')) as avg_approval_time,
    (SELECT COALESCE((COUNT(*) FILTER (WHERE decided_at - created_at <= (SELECT COALESCE(NULLIF(value, ''), '24') FROM system_settings WHERE key = 'sla_approval_hours')::int * interval '1 hour')::float / NULLIF(COUNT(*), 0)::float) * 100, 100)::float FROM approval_workflows WHERE approver_id = $1 AND status IN ('approved', 'rejected')) as compliance_rate;

-- name: GetManagerTopSubmitters :many
SELECT 
    u.full_name as name,
    u.avatar_url as avatar,
    COUNT(d.id) as count
FROM users u
LEFT JOIN documents d ON u.id = d.owner_id
WHERE u.department_id IN (SELECT id FROM departments WHERE head_id = $1)
GROUP BY u.id, u.full_name, u.avatar_url
ORDER BY count DESC
LIMIT $2;

-- name: GetManagerRecentActivities :many
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
WHERE u.department_id IN (SELECT id FROM departments WHERE head_id = $1)
ORDER BY a.created_at DESC
LIMIT $2;
-- name: ApproveTask :exec
UPDATE approval_workflows 
SET status = 'approved', decided_at = NOW(), decision_note = $2
WHERE entity_id = $1 AND status = 'pending';

-- name: RejectTask :exec
UPDATE approval_workflows 
SET status = 'rejected', decided_at = NOW(), decision_note = $2, rejection_reason = $3
WHERE entity_id = $1 AND status = 'pending';

-- name: GetPendingCountsByType :many
SELECT entity_type, COUNT(*) as count
FROM approval_workflows
WHERE approver_id = $1 AND status = 'pending'
GROUP BY entity_type;

-- name: CreateApprovalTask :one
INSERT INTO approval_workflows (
    entity_type, entity_id, approver_id, level, status
) VALUES (
    $1, $2, $3, $4, 'pending'
) RETURNING *;

-- name: GetLatestApprovalTaskByEntity :one
SELECT * FROM approval_workflows
WHERE entity_id = $1 AND entity_type = $2
ORDER BY created_at DESC
LIMIT 1;
