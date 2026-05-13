-- 000071_add_intake_manager_notification_template.up.sql

INSERT INTO email_templates (slug, name, subject, body_html, placeholders) VALUES 
(
    'intake-received-manager', 
    'Physical Document Received (Manager Notification)', 
    'Kreatif DMS - Penerimaan Berkas Departemen: {{manifestNo}}', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #1E3A5F;">Penerimaan Berkas Fisik</h2>
        <p>Halo <strong>{{fullName}}</strong>,</p>
        <p>Kami informasikan bahwa berkas fisik untuk manifest <strong>{{manifestNo}}</strong> ({{itemCount}} dokumen) dari <strong>{{senderName}}</strong> telah diterima secara lengkap oleh tim Central Document.</p>
        <div style="background: #f0f7ff; padding: 15px; border-radius: 8px; border-left: 4px solid #1E3A5F; margin: 20px 0;">
            <p style="margin: 0;"><strong>Status:</strong> Aktif & Terarsip</p>
            <p style="margin: 5px 0 0 0; color: #666;">Seluruh dokumen digital dalam manifest ini sekarang telah aktif sepenuhnya di dalam sistem.</p>
        </div>
        <p>Salam hangat,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName", "manifestNo", "itemCount", "senderName"]'
) ON CONFLICT (slug) DO NOTHING;
