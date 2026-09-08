-- name: CreateLedgerEntry :one
INSERT INTO ledger_entries (
  account_id, transaction_id, credit_debit, amount, balance_after
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, account_id, transaction_id, credit_debit, amount, balance_after, created_at;

-- name: ListLedgerEntriesByAccountId :many
SELECT id, account_id, transaction_id, credit_debit, amount, balance_after, created_at
FROM ledger_entries
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;