-- 000079_add_penalty_policy_setting.up.sql

INSERT INTO system_settings (category, key, value, value_type, description)
VALUES (
    'general',
    'penalty_policy',
    'Denda keterlambatan: IDR 50.000 / Dokumen / Hari',
    'string',
    'Kebijakan denda/penalti keterlambatan pengembalian dokumen'
) ON CONFLICT (category, key) DO UPDATE SET value = EXCLUDED.value;
