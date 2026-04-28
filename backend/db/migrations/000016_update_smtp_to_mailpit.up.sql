-- Migration: Update SMTP config to Mailpit
UPDATE integration_nodes 
SET 
    name = 'Mailpit (Local Dev)',
    driver = 'mailpit',
    endpoint = '127.0.0.1:1025',
    config_json = '{"from_email": "dms@kreatif.id", "auth": false}'
WHERE service_type = 'SMTP';
