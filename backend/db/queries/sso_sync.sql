-- name: ListSsoSyncLogs :many
SELECT * FROM sso_sync_logs
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetLastSsoSyncLog :one
SELECT * FROM sso_sync_logs
ORDER BY created_at DESC
LIMIT 1;

-- name: CreateSsoSyncLog :one
INSERT INTO sso_sync_logs (
    provider, status, started_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateSsoSyncLog :one
UPDATE sso_sync_logs
SET 
    status = $2,
    users_synced = $3,
    groups_synced = $4,
    errors = $5,
    error_details = $6,
    completed_at = $7
WHERE id = $1
RETURNING *;

-- name: CountSsoSyncLogs :one
SELECT COUNT(*) FROM sso_sync_logs;
