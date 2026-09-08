-- name: CreateAuditLog :one
INSERT INTO audit_logs (
  actor_id, action, table_name, old_values, new_values, ip_address
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, actor_id, action, table_name, old_values, new_values, ip_address, created_at;

-- name: ListAuditLogsByActorId :many
SELECT id, actor_id, action, table_name, old_values, new_values, ip_address, created_at
FROM audit_logs
WHERE actor_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;