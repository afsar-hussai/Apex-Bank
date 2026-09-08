-- name: CreateAccount :one
INSERT INTO accounts (
  user_id, balance, currency, account_type
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, user_id, balance, currency, account_type, created_at, updated_at;

-- name: GetAccount :one
SELECT id, user_id, balance, currency, account_type, created_at, updated_at
FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccountsByUserId :many
SELECT id, user_id, balance, currency, account_type, created_at, updated_at
FROM accounts
WHERE user_id = $1
ORDER BY id;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount), updated_at = now()
WHERE id = sqlc.arg(account_id)
RETURNING id, user_id, balance, currency, account_type, created_at, updated_at;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;