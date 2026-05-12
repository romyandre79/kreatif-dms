-- name: ListNotifications :many
SELECT * FROM notifications 
WHERE user_id = $1 
ORDER BY created_at DESC 
LIMIT $2 OFFSET $3;

-- name: GetUnreadCount :one
SELECT COUNT(*) FROM notifications 
WHERE user_id = $1 AND is_read = false;

-- name: MarkAsRead :exec
UPDATE notifications 
SET is_read = true, read_at = NOW() 
WHERE id = $1 AND user_id = $2;

-- name: MarkAllAsRead :exec
UPDATE notifications 
SET is_read = true, read_at = NOW() 
WHERE user_id = $1 AND is_read = false;

-- name: CreateNotification :one
INSERT INTO notifications (
    user_id, title, body, type, entity_type, entity_id, channel, metadata
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;
