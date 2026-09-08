-- name: CreateTransaction :one
INSERT INTO transactions (
  sender_account_id, receiver_account_id, amount, currency, status, transaction_type, idempotency_key, description
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, sender_account_id, receiver_account_id, amount, currency, status, transaction_type, idempotency_key, description, created_at;

-- name: GetTransactionById :one
SELECT id, sender_account_id, receiver_account_id, amount, currency, status, transaction_type, idempotency_key, description, created_at
FROM transactions
WHERE id = $1 LIMIT 1;

-- name: GetTransactionByIdempotencyKey :one
SELECT id, sender_account_id, receiver_account_id, amount, currency, status, transaction_type, idempotency_key, description, created_at
FROM transactions
WHERE idempotency_key = $1 LIMIT 1;

-- name: ListTransactionsByAccountId :many
SELECT id, sender_account_id, receiver_account_id, amount, currency, status, transaction_type, idempotency_key, description, created_at
FROM transactions
WHERE sender_account_id = $1 OR receiver_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;