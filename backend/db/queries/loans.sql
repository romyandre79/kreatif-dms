-- name: CreateLoanRequest :one
INSERT INTO loan_requests (
    request_no, user_id, purpose, duration_days, notes, status,
    company_id, branch_id, department_id
) VALUES (
    $1, $2, $3, $4, $5, 'pending', $6, $7, $8
) RETURNING id, request_no, user_id, purpose, department_filter, duration_days, notes, approver_notes, status, l1_approved_by, l1_approved_at, l1_rejection_reason, l2_approved_by, l2_approved_at, l2_rejection_reason, borrow_date, due_date, return_date, created_at, updated_at, company_id, branch_id, department_id;

-- name: CreateLoanRequestItem :one
INSERT INTO loan_request_items (
    loan_request_id, document_id, category, sensitivity, method
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, loan_request_id, document_id, category, sensitivity, reason, method, status, picked_up_at, returned_at, created_at;

-- name: GetLoanRequest :one
SELECT 
    lr.id, lr.request_no, lr.user_id, lr.purpose, lr.duration_days, lr.notes, 
    lr.status, lr.borrow_date, lr.due_date, lr.return_date,
    lr.l1_approved_by, lr.l1_approved_at, lr.l2_approved_by, lr.l2_approved_at,
    lr.created_at, lr.updated_at,
    u.full_name as user_name,
    dept.name as department_name,
    (SELECT COUNT(*) FROM loan_request_items WHERE loan_request_id = lr.id) as items_count
FROM loan_requests lr
JOIN users u ON lr.user_id = u.id
LEFT JOIN departments dept ON lr.department_id = dept.id
WHERE lr.id = $1 LIMIT 1;

-- name: ListUserLoanRequests :many
SELECT 
    lr.id, lr.request_no, lr.user_id, lr.purpose, lr.duration_days, lr.notes, 
    lr.status, lr.borrow_date, lr.due_date, lr.return_date,
    lr.l1_approved_by, lr.l1_approved_at,
    lr.created_at, lr.updated_at,
    u.full_name as user_name,
    dept.name as department_name,
    (SELECT COUNT(*) FROM loan_request_items WHERE loan_request_id = lr.id) as items_count
FROM loan_requests lr
JOIN users u ON lr.user_id = u.id
LEFT JOIN departments dept ON lr.department_id = dept.id
WHERE lr.user_id = $1
ORDER BY lr.created_at DESC;

-- name: ListAllLoanRequests :many
SELECT 
    lr.id, lr.request_no, lr.user_id, lr.purpose, lr.duration_days, lr.notes, 
    lr.status, lr.borrow_date, lr.due_date, lr.return_date,
    lr.l1_approved_by, lr.l1_approved_at,
    lr.created_at, lr.updated_at,
    u.full_name as user_name,
    dept.name as department_name,
    (SELECT COUNT(*) FROM loan_request_items WHERE loan_request_id = lr.id) as items_count
FROM loan_requests lr
JOIN users u ON lr.user_id = u.id
LEFT JOIN departments dept ON lr.department_id = dept.id
ORDER BY lr.created_at DESC
LIMIT $1;

-- name: GetLoanRequestItems :many
SELECT 
    lri.id, lri.loan_request_id, lri.document_id, lri.category, lri.sensitivity,
    lri.method, lri.status, lri.picked_up_at, lri.returned_at, lri.created_at,
    d.title as document_title,
    d.file_name as document_filename
FROM loan_request_items lri
JOIN documents d ON lri.document_id = d.id
WHERE lri.loan_request_id = $1
ORDER BY lri.created_at ASC;

-- name: UpdateLoanRequestStatus :exec
UPDATE loan_requests 
SET status = $2, updated_at = NOW()
WHERE id = $1;

-- name: ApproveLoanRequestL1 :exec
UPDATE loan_requests
SET status = 'l1_approved', l1_approved_by = $2, l1_approved_at = NOW(), 
    borrow_date = NOW(), due_date = NOW() + (duration_days || ' days')::interval,
    updated_at = NOW()
WHERE id = $1;

-- name: RejectLoanRequest :exec
UPDATE loan_requests
SET status = 'rejected', l1_approved_by = $2, l1_rejection_reason = $3, updated_at = NOW()
WHERE id = $1;

-- name: CountUserLoanRequests :one
SELECT COUNT(*) FROM loan_requests WHERE user_id = $1;

-- name: GetNextLoanRequestNo :one
SELECT COALESCE(MAX(CAST(SUBSTRING(request_no FROM 'LOAN-\d{4}-\d{2}-(\d+)') AS INT)), 0) + 1 as next_seq
FROM loan_requests
WHERE request_no LIKE 'LOAN-' || to_char(NOW(), 'YYYY-MM') || '-%';
