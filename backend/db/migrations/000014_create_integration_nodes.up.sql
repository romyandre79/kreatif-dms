-- Function to update updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Migration: Add integration_nodes table
CREATE TABLE IF NOT EXISTS integration_nodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    service_type VARCHAR(50) NOT NULL, -- 'LDAP', 'S3', 'DATABASE', 'SMTP', 'SCANNER', 'PRINTER', 'RFID'
    driver VARCHAR(50),                -- 'openldap', 'minio', 'postgres', 'office365', 'zpl'
    endpoint VARCHAR(255) NOT NULL,    -- Host/IP:Port
    is_active BOOLEAN DEFAULT true,
    is_critical BOOLEAN DEFAULT false,
    config_json JSONB DEFAULT '{}',     -- Kredensial & setting spesifik (encrypted if needed)
    
    -- Telemetry Data
    status VARCHAR(20) DEFAULT 'unknown', -- 'online', 'offline', 'warning', 'unknown'
    last_latency INT DEFAULT 0,           -- ms
    latency_history INT[] DEFAULT '{}',    -- 10 data terakhir
    last_check_at TIMESTAMP WITH TIME ZONE,
    last_error TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger for updated_at
DROP TRIGGER IF EXISTS update_integration_nodes_updated_at ON integration_nodes;
CREATE TRIGGER update_integration_nodes_updated_at
    BEFORE UPDATE ON integration_nodes
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Seed Initial Data from .env (Manual values for now based on previous context)
-- Note: Replace these with actual values if you want me to do it now.
INSERT INTO integration_nodes (name, service_type, driver, endpoint, is_critical, config_json)
VALUES 
('Main LDAP Server', 'LDAP', 'openldap', '127.0.0.1:3890', true, '{"base_dn": "dc=example,dc=org", "bind_dn": "cn=admin,dc=example,dc=org"}'),
('Primary S3 Storage', 'S3', 'minio', '127.0.0.1:9000', true, '{"bucket": "kreatif-dms", "use_ssl": false}'),
('Corporate SMTP Relay', 'SMTP', 'office365', 'smtp.office365.com:587', false, '{"from_email": "noreply@kreatif.com"}');
