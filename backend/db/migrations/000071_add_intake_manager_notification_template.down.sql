-- 000071_add_intake_manager_notification_template.down.sql

DELETE FROM email_templates WHERE slug = 'intake-received-manager';
