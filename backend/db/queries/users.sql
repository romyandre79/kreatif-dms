-- name: GetUserByEmail :one
SELECT u.id, u.email, u.password_hash, u.full_name, u.role_id, u.department_id, u.is_active, u.created_at, u.updated_at, u.status, u.avatar_url, u.signature_url, r.name as role_name 
FROM users u
LEFT JOIN roles r ON u.role_id = r.id
WHERE u.email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT id, email, password_hash, full_name, role_id, department_id, is_active, created_at, updated_at, status, avatar_url, signature_url FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    email, password_hash, full_name, role_id, department_id, status
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, email, password_hash, full_name, role_id, department_id, is_active, created_at, updated_at, status, avatar_url, signature_url;

-- name: UpdateUserStatus :one
UPDATE users 
SET status = $2, updated_at = NOW()
WHERE id = $1
RETURNING id, email, password_hash, full_name, role_id, department_id, is_active, created_at, updated_at, status, avatar_url, signature_url;

-- name: ListUsers :many
SELECT id, email, password_hash, full_name, role_id, department_id, is_active, created_at, updated_at, status, avatar_url, signature_url FROM users ORDER BY created_at DESC;

-- name: ListPendingUsers :many
SELECT id, email, password_hash, full_name, role_id, department_id, is_active, created_at, updated_at, status, avatar_url, signature_url FROM users WHERE status = 'pending' ORDER BY created_at ASC;

-- name: GetRoleIDByName :one
SELECT id FROM roles WHERE name = $1 LIMIT 1;

-- name: UpdateUserPIN :exec
UPDATE users SET pin = $2 WHERE id = $1;
