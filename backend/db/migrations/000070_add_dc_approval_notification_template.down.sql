-- 000070_add_dc_approval_notification_template.down.sql

DELETE FROM email_templates WHERE slug = 'doc-approved-dc';
