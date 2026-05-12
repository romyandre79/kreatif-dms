-- name: ListEmailTemplates :many
SELECT * FROM email_templates ORDER BY name ASC;

-- name: GetEmailTemplateBySlug :one
SELECT * FROM email_templates WHERE slug = $1 LIMIT 1;

-- name: UpdateEmailTemplate :one
UPDATE email_templates 
SET subject = $2, body_html = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;
