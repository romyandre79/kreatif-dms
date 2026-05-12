-- Companies
-- name: ListCompanies :many
SELECT * FROM companies ORDER BY name;

-- name: GetCompany :one
SELECT * FROM companies WHERE id = $1;

-- name: CreateCompany :one
INSERT INTO companies (name, entity_id, npwp_status, location, status, address, delivery_instructions)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateCompany :one
UPDATE companies SET 
    name = $2, 
    entity_id = $3, 
    npwp_status = $4, 
    location = $5, 
    status = $6, 
    address = $7,
    delivery_instructions = $8
WHERE id = $1
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM companies WHERE id = $1;

-- name: UpdateCompanyLogo :exec
UPDATE companies SET logo_url = $2 WHERE id = $1;

-- Branches
-- name: ListBranches :many
SELECT * FROM branches WHERE company_id = $1 ORDER BY name;

-- name: ListAllBranchesGlobal :many
SELECT 
    b.*, 
    c.name as company_name,
    u.full_name as head_name
FROM branches b
JOIN companies c ON b.company_id = c.id
LEFT JOIN users u ON b.head_id = u.id
ORDER BY b.name;

-- name: GetBranch :one
SELECT * FROM branches WHERE id = $1;

-- name: CreateBranch :one
INSERT INTO branches (company_id, name, location, head_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateBranch :one
UPDATE branches SET name = $2, location = $3, head_id = $4, company_id = $5
WHERE id = $1
RETURNING *;

-- name: DeleteBranch :exec
DELETE FROM branches WHERE id = $1;

-- Departments
-- name: ListDepartments :many
SELECT * FROM departments WHERE branch_id = $1 ORDER BY name;

-- name: ListAllDepartments :many
SELECT 
    d.*, 
    b.name as branch_name,
    u.full_name as head_name
FROM departments d
JOIN branches b ON d.branch_id = b.id
LEFT JOIN users u ON d.head_id = u.id
ORDER BY d.name;

-- name: GetDepartment :one
SELECT * FROM departments WHERE id = $1;

-- name: CreateDepartment :one
INSERT INTO departments (branch_id, name, head_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateDepartment :one
UPDATE departments SET name = $2, head_id = $3, branch_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteDepartment :exec
DELETE FROM departments WHERE id = $1;

-- Racks
-- name: ListRacks :many
SELECT * FROM racks WHERE department_id = $1 ORDER BY name;

-- name: ListAllRacksGlobal :many
SELECT r.*, d.name as department_name, b.name as branch_name
FROM racks r
JOIN departments d ON r.department_id = d.id
JOIN branches b ON d.branch_id = b.id
ORDER BY r.name;

-- name: GetRack :one
SELECT * FROM racks WHERE id = $1;

-- name: CreateRack :one
INSERT INTO racks (department_id, name, location_detail)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateRack :one
UPDATE racks SET name = $2, location_detail = $3, department_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteRack :exec
DELETE FROM racks WHERE id = $1;

-- Boxes
-- name: ListBoxes :many
SELECT * FROM boxes WHERE rack_id = $1 ORDER BY name;

-- name: ListAllBoxesGlobal :many
SELECT bx.*, r.name as rack_name, d.name as department_name
FROM boxes bx
JOIN racks r ON bx.rack_id = r.id
JOIN departments d ON r.department_id = d.id
ORDER BY bx.name;

-- name: GetBox :one
SELECT * FROM boxes WHERE id = $1;

-- name: CreateBox :one
INSERT INTO boxes (rack_id, name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateBox :one
UPDATE boxes SET name = $2, rack_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBox :exec
DELETE FROM boxes WHERE id = $1;

-- Ordners
-- name: ListOrdners :many
SELECT * FROM ordners WHERE box_id = $1 ORDER BY name;

-- name: ListAllOrdnersGlobal :many
SELECT o.*, bx.name as box_name, r.name as rack_name
FROM ordners o
JOIN boxes bx ON o.box_id = bx.id
JOIN racks r ON bx.rack_id = r.id
ORDER BY o.name;
-- Roles
-- name: ListRoles :many
SELECT r.*, (SELECT COUNT(*) FROM users u WHERE u.role_id = r.id) as user_count
FROM roles r ORDER BY r.name;

-- name: GetRole :one
SELECT * FROM roles WHERE id = $1;

-- name: CreateRole :one
INSERT INTO roles (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateRole :one
UPDATE roles SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1;

-- System Modules & Permissions
-- name: ListSystemModules :many
SELECT * FROM system_modules ORDER BY COALESCE(parent_id, ''), sort_order, name;

-- name: GetSystemModule :one
SELECT * FROM system_modules WHERE id = $1;

-- name: CreateSystemModule :one
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateSystemModule :one
UPDATE system_modules SET 
    name = $2, 
    category = $3, 
    path = $4, 
    icon = $5, 
    allowed_actions = $6, 
    sort_order = $7,
    parent_id = $8
WHERE id = $1
RETURNING *;

-- name: DeleteSystemModule :exec
DELETE FROM system_modules WHERE id = $1;

-- name: GetRolePermissions :many
SELECT module_id, action FROM role_permissions WHERE role_id = $1;

-- name: ClearRolePermissions :exec
DELETE FROM role_permissions WHERE role_id = $1;

-- name: AddRolePermission :exec
INSERT INTO role_permissions (role_id, module_id, action) VALUES ($1, $2, $3);

-- name: ListPermissionsByRole :many
SELECT m.id as module_id, m.name as module_name, m.category, rp.action
FROM role_permissions rp
JOIN system_modules m ON rp.module_id = m.id
WHERE rp.role_id = $1;

-- RFID Tags
-- name: ListRfidTags :many
SELECT * FROM rfid_tags ORDER BY assigned_at DESC;

-- name: GetRfidTag :one
SELECT * FROM rfid_tags WHERE tag_id = $1;

-- name: CreateRfidTag :one
INSERT INTO rfid_tags (tag_id, document_id, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateRfidTagStatus :one
UPDATE rfid_tags SET status = $2
WHERE tag_id = $1
RETURNING *;

-- name: GetOrdner :one
SELECT * FROM ordners WHERE id = $1;

-- name: CreateOrdner :one
INSERT INTO ordners (box_id, name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateOrdner :one
UPDATE ordners SET name = $2, box_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteOrdner :exec
DELETE FROM ordners WHERE id = $1;

-- name: GetWarehouseTopology :many
SELECT 
    c.id as company_id, c.name as company_name, 
    b.id as branch_id, b.name as branch_name, 
    d.id as department_id, d.name as department_name,
    r.id as rack_id, r.name as rack_name,
    bx.id as box_id, bx.name as box_name,
    o.id as ordner_id, o.name as ordner_name
FROM companies c
JOIN branches b ON b.company_id = c.id
JOIN departments d ON d.branch_id = b.id
JOIN racks r ON r.department_id = d.id
LEFT JOIN boxes bx ON bx.rack_id = r.id
LEFT JOIN ordners o ON o.box_id = bx.id
ORDER BY c.name, b.name, d.name, r.name, bx.name, o.name;

-- Retention Policies
-- name: ListRetentionPolicies :many
SELECT * FROM retention_policies ORDER BY name;

-- name: CreateRetentionPolicy :one
INSERT INTO retention_policies (name, description, retention_years, department_id, document_category)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateRetentionPolicy :one
UPDATE retention_policies SET name = $2, description = $3, retention_years = $4, department_id = $5, document_category = $6
WHERE id = $1
RETURNING *;

-- name: DeleteRetentionPolicy :exec
DELETE FROM retention_policies WHERE id = $1;

-- System Settings
-- name: GetSystemSettingsByCategory :many
SELECT * FROM system_settings WHERE category = $1;

-- name: GetSystemSetting :one
SELECT * FROM system_settings WHERE category = $1 AND key = $2;

-- name: UpsertSystemSetting :one
INSERT INTO system_settings (category, key, value, value_type, description, updated_by)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (category, key) DO UPDATE
SET value = EXCLUDED.value, value_type = EXCLUDED.value_type, description = EXCLUDED.description, updated_by = EXCLUDED.updated_by, updated_at = NOW()
RETURNING *;

-- Document Categories
-- name: ListDocumentCategories :many
SELECT * FROM document_categories ORDER BY name;

-- name: GetDocumentCategory :one
SELECT * FROM document_categories WHERE id = $1;

-- name: CreateDocumentCategory :one
INSERT INTO document_categories (code, name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- Document Types
-- name: ListDocumentTypes :many
SELECT 
    dt.id, dt.code, dt.name, dt.description, dt.created_at, dt.updated_at, dt.category_id,
    dc.name as category_name 
FROM document_types dt
LEFT JOIN document_categories dc ON dt.category_id = dc.id
ORDER BY dt.name;

-- name: GetDocumentType :one
SELECT * FROM document_types WHERE id = $1;

-- name: CreateDocumentType :one
INSERT INTO document_types (code, name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateDocumentType :one
UPDATE document_types SET code = $2, name = $3, description = $4, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteDocumentType :exec
DELETE FROM document_types WHERE id = $1;
