-- 000079_add_penalty_policy_setting.down.sql

DELETE FROM system_settings WHERE category = 'general' AND key = 'penalty_policy';
