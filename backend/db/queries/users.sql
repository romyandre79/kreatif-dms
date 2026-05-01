-- name: GetUserByEmail :one
SELECT u.*, r.name as role_name 
FROM users u
LEFT JOIN roles r ON u.role_id = r.id
WHERE u.email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT u.*, r.name as role_name 
FROM users u
LEFT JOIN roles r ON u.role_id = r.id
WHERE u.id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    email, password_hash, full_name, role_id, department_id, status, avatar_url, signature_url, is_mfa_enabled
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: UpdateUserStatus :one
UPDATE users 
SET status = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: ListUsers :many
SELECT u.*, r.name as role_name 
FROM users u
LEFT JOIN roles r ON u.role_id = r.id
ORDER BY u.created_at DESC;

-- name: ListPendingUsers :many
SELECT * FROM users WHERE status = 'pending' ORDER BY created_at ASC;

-- name: GetRoleIDByName :one
SELECT id FROM roles WHERE name = $1 LIMIT 1;

-- name: UpdateUserPIN :exec
UPDATE users SET pin = $2 WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET email = $2, full_name = $3, role_id = $4, department_id = $5, status = $6, 
    avatar_url = $7, signature_url = $8, is_mfa_enabled = $9, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUserMFASecret :exec
UPDATE users SET mfa_secret = $2, is_mfa_enabled = $3 WHERE id = $1;
