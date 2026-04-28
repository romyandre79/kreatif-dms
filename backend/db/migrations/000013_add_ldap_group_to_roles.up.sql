-- Add ldap_group column to roles table
ALTER TABLE roles ADD COLUMN IF NOT EXISTS ldap_group VARCHAR(100);

-- Update existing roles with default mapping for LLDAP
UPDATE roles SET ldap_group = 'lldap_admin' WHERE name = 'admin';
UPDATE roles SET ldap_group = 'lldap_manager' WHERE name = 'manager';
UPDATE roles SET ldap_group = 'lldap_user' WHERE name = 'user';
