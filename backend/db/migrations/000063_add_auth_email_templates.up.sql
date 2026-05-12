-- 000063_add_auth_email_templates.up.sql

INSERT INTO email_templates (slug, name, subject, body_html, placeholders) VALUES 
(
    'password-reset', 
    'Password Reset Request', 
    'Kreatif DMS - Reset Your Password', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #1E3A5F;">Reset Your Password</h2>
        <p>We received a request to reset your password for your Kreatif DMS account.</p>
        <p>Click the button below to reset it:</p>
        <div style="text-align: center; margin: 30px 0;">
            <a href="{{resetLink}}" style="background: #1E3A5F; color: white; padding: 12px 25px; text-decoration: none; border-radius: 5px; font-weight: bold; display: inline-block;">Reset Password</a>
        </div>
        <p style="color: #666; font-size: 12px;">If you didn''t request this, you can safely ignore this email. The link will expire in 1 hour.</p>
        <p>Best regards,<br>Kreatif DMS Team</p>
    </div>',
    '["resetLink"]'
),
(
    'registration-success', 
    'Registration Successful', 
    'Kreatif DMS - Welcome to the Platform', 
    '<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
        <h2 style="color: #1E3A5F;">Welcome to Kreatif DMS</h2>
        <p>Hi <strong>{{fullName}}</strong>,</p>
        <p>Thank you for registering at Kreatif DMS.</p>
        <div style="background: #f0f7ff; padding: 15px; border-radius: 8px; border-left: 4px solid #1E3A5F; margin: 20px 0;">
            <p style="margin: 0;"><strong>Account Status: Pending Approval</strong></p>
            <p style="margin: 5px 0 0 0; color: #666;">Your account is currently being reviewed by our administrator. You will receive another email once your account is approved.</p>
        </div>
        <p>Best regards,<br>Kreatif DMS Team</p>
    </div>',
    '["fullName"]'
)
ON CONFLICT (slug) DO NOTHING;
