-- 000062_create_email_templates.up.sql

CREATE TABLE IF NOT EXISTS email_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    subject VARCHAR(500) NOT NULL,
    body_html TEXT NOT NULL,
    placeholders JSONB DEFAULT '[]',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Seed initial templates
INSERT INTO email_templates (slug, name, subject, body_html, placeholders) VALUES 
(
    'doc-approved', 
    'Document Approved', 
    'Kreatif DMS - Document Approved: {{docTitle}}', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #22c55e;">Document Approved!</h2>
        <p>Hi <strong>{{fullName}}</strong>,</p>
        <p>Your document "<strong>{{docTitle}}</strong>" has been approved.</p>
        <div style="background: #f0fdf4; padding: 15px; border-radius: 8px; border-left: 4px solid #22c55e; margin: 20px 0;">
            <p style="margin: 0;"><strong>Notes from Approver:</strong></p>
            <p style="margin: 5px 0 0 0; color: #666;">{{notes}}</p>
        </div>
        <p>You can now view the document in the repository.</p>
        <p>Best regards,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName", "docTitle", "notes"]'
),
(
    'doc-rejected', 
    'Document Rejected', 
    'Kreatif DMS - Document Rejected: {{docTitle}}', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #ef4444;">Document Rejected</h2>
        <p>Hi <strong>{{fullName}}</strong>,</p>
        <p>We regret to inform you that your document "<strong>{{docTitle}}</strong>" has been rejected.</p>
        <div style="background: #fef2f2; padding: 15px; border-radius: 8px; border-left: 4px solid #ef4444; margin: 20px 0;">
            <p style="margin: 0;"><strong>Reason for Rejection:</strong></p>
            <p style="margin: 5px 0 0 0; color: #666;">{{notes}}</p>
        </div>
        <p>Please review the notes and resubmit if necessary.</p>
        <p>Best regards,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName", "docTitle", "notes"]'
),
(
    'intake-received', 
    'Physical Document Received', 
    'Kreatif DMS - Physical Receipt: {{manifestNo}}', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #1E3A5F;">Physical Documents Received</h2>
        <p>Hi <strong>{{fullName}}</strong>,</p>
        <p>Your physical documents for manifest <strong>{{manifestNo}}</strong> ({{itemCount}} items) have been successfully received and verified by the warehouse team.</p>
        <div style="background: #f0f7ff; padding: 15px; border-radius: 8px; border-left: 4px solid #1E3A5F; margin: 20px 0;">
            <p style="margin: 0;"><strong>Status:</strong> Archived & Active</p>
            <p style="margin: 5px 0 0 0; color: #666;">The documents are now physically stored and their digital counterparts are fully activated in the system.</p>
        </div>
        <p>Best regards,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName", "manifestNo", "itemCount"]'
),
(
    'intake-rejected', 
    'Physical Document Rejected', 
    'Kreatif DMS - Physical Rejection: {{manifestNo}}', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #f59e0b;">Physical Documents Rejected</h2>
        <p>Hi <strong>{{fullName}}</strong>,</p>
        <p>There was an issue with the physical delivery of manifest <strong>{{manifestNo}}</strong>.</p>
        <div style="background: #fffbeb; padding: 15px; border-radius: 8px; border-left: 4px solid #f59e0b; margin: 20px 0;">
            <p style="margin: 0;"><strong>Reason:</strong></p>
            <p style="margin: 5px 0 0 0; color: #666;">{{reason}}</p>
        </div>
        <p>Please contact the warehouse department for further details.</p>
        <p>Best regards,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName", "manifestNo", "reason"]'
);
