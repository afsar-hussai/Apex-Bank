-- name: CreateUser :one
INSERT INTO users (
  full_name, email, phone, password_hash, role, kyc_status
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, full_name, email, phone, password_hash, role, kyc_status, is_active, created_at, updated_at;

-- name: GetUserById :one
SELECT id, full_name, email, phone, password_hash, role, kyc_status, is_active, created_at, updated_at
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, full_name, email, phone, password_hash, role, kyc_status, is_active, created_at, updated_at
FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, full_name, email, phone, password_hash, role, kyc_status, is_active, created_at, updated_at
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateUserKycStatus :one
UPDATE users
SET kyc_status = $2, updated_at = now()
WHERE id = $1
RETURNING id, full_name, email, phone, password_hash, role, kyc_status, is_active, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;