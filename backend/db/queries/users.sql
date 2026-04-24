-- name: GetUserByEmail :one
SELECT u.*, r.name as role_name 
FROM users u
JOIN roles r ON u.role_id = r.id
WHERE u.email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    email, password_hash, full_name, role_id, department_id, status
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateUserStatus :one
UPDATE users 
SET status = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: ListPendingUsers :many
SELECT * FROM users WHERE status = 'pending' ORDER BY created_at ASC;

-- name: GetRoleIDByName :one
SELECT id FROM roles WHERE name = $1 LIMIT 1;
