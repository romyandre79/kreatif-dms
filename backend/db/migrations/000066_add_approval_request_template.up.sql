-- 000066_add_approval_request_template.up.sql

INSERT INTO email_templates (slug, name, subject, body_html, placeholders) VALUES 
(
    'doc-pending-approval', 
    'Approval Request', 
    'Kreatif DMS - Permintaan Persetujuan: {{docTitle}}', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #1E3A5F;">Permintaan Persetujuan</h2>
        <p>Halo <strong>{{fullName}}</strong>,</p>
        <p>Dokumen baru "<strong>{{docTitle}}</strong>" dari departemen <strong>{{departmentName}}</strong> memerlukan persetujuan Anda.</p>
        <div style="background: #f0f7ff; padding: 15px; border-radius: 8px; border-left: 4px solid #1E3A5F; margin: 20px 0;">
            <p style="margin: 0;"><strong>Pengunggah:</strong> {{ownerName}}</p>
            <p style="margin: 5px 0 0 0; color: #666;">Silakan tinjau dokumen ini melalui tautan di bawah ini.</p>
        </div>
        <div style="text-align: center; margin: 30px 0;">
            <a href="{{approveLink}}" style="background: #1E3A5F; color: white; padding: 12px 25px; text-decoration: none; border-radius: 5px; font-weight: bold; display: inline-block;">Lihat & Setujui Dokumen</a>
        </div>
        <p style="font-size: 12px; color: #999;">Jika tombol di atas tidak berfungsi, Anda juga bisa menyalin tautan berikut ke browser Anda:<br>{{approveLink}}</p>
        <p>Salam hangat,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName", "docTitle", "departmentName", "ownerName", "approveLink"]'
) ON CONFLICT (slug) DO NOTHING;
