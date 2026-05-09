-- Announcements
-- name: ListAnnouncements :many
SELECT * FROM announcements ORDER BY created_at DESC;

-- name: GetActiveAnnouncement :one
SELECT * FROM announcements WHERE is_active = true ORDER BY created_at DESC LIMIT 1;

-- name: GetAnnouncement :one
SELECT * FROM announcements WHERE id = $1;

-- name: CreateAnnouncement :one
INSERT INTO announcements (title, message, notes, is_active, created_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateAnnouncement :one
UPDATE announcements SET 
    title = $2, 
    message = $3, 
    notes = $4, 
    is_active = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAnnouncement :exec
DELETE FROM announcements WHERE id = $1;

-- name: DeactivateAllAnnouncements :exec
UPDATE announcements SET is_active = false WHERE is_active = true;
