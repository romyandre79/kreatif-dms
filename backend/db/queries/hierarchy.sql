-- name: GetCompanies :many
SELECT id, name, address, created_at FROM companies ORDER BY name ASC;

-- name: GetBranches :many
SELECT id, company_id, name, location, head_id FROM branches ORDER BY name ASC;

-- name: GetDepartments :many
SELECT id, branch_id, name, head_id FROM departments ORDER BY name ASC;

-- name: GetUsersForHierarchy :many
SELECT u.id, u.full_name, u.role_id, r.name as role_name, u.department_id
FROM users u
LEFT JOIN roles r ON u.role_id = r.id
WHERE u.status = 'approved' OR u.status = 'active'
ORDER BY u.full_name ASC;

-- name: UpdateUserDepartment :exec
UPDATE users SET department_id = $1 WHERE id = $2;

-- name: UpdateDepartmentHead :exec
UPDATE departments SET head_id = $1 WHERE id = $2;