-- name: ListIntegrationNodes :many
SELECT * FROM integration_nodes
ORDER BY service_type, name;

-- name: GetIntegrationNode :one
SELECT * FROM integration_nodes
WHERE id = $1 LIMIT 1;

-- name: GetIntegrationNodeByType :one
SELECT * FROM integration_nodes
WHERE service_type = $1 AND is_active = true
LIMIT 1;

-- name: CreateIntegrationNode :one
INSERT INTO integration_nodes (
    name, service_type, driver, endpoint, is_critical, config_json
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateIntegrationNodeStatus :one
UPDATE integration_nodes
SET 
    status = $2,
    last_latency = $3,
    latency_history = array_append(latency_history, $3::int),
    last_check_at = CURRENT_TIMESTAMP,
    last_error = $4
WHERE id = $1
RETURNING *;

-- name: UpdateIntegrationNodeConfig :one
UPDATE integration_nodes
SET 
    name = $2,
    endpoint = $3,
    is_active = $4,
    is_critical = $5,
    config_json = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteIntegrationNode :exec
DELETE FROM integration_nodes
WHERE id = $1;
