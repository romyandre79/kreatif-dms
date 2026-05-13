-- 000070_add_dc_approval_notification_template.up.sql

INSERT INTO email_templates (slug, name, subject, body_html, placeholders) VALUES 
(
    'doc-approved-dc', 
    'Document Approved (DC Notification)', 
    'Kreatif DMS - Dokumen Siap Intake: {{docTitle}}', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #1E3A5F;">Dokumen Siap Intake</h2>
        <p>Halo <strong>{{fullName}}</strong>,</p>
        <p>Dokumen "<strong>{{docTitle}}</strong>" telah disetujui oleh Manager dan sekarang menunggu proses <strong>Physical Intake</strong>.</p>
        <div style="background: #f0f7ff; padding: 15px; border-radius: 8px; border-left: 4px solid #1E3A5F; margin: 20px 0;">
            <p style="margin: 0;"><strong>Pengunggah:</strong> {{ownerName}}</p>
            <p style="margin: 5px 0 0 0; color: #666;">Silakan periksa antrean di dashboard Document Controller untuk memproses dokumen ini.</p>
        </div>
        <p>Salam hangat,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName", "docTitle", "ownerName"]'
) ON CONFLICT (slug) DO NOTHING;
