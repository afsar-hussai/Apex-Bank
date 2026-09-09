package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLedgerEntry(t *testing.T) {
	account := createRandomAccount(t)
	tx := createRandomTransaction(t)

	arg := CreateLedgerEntryParams{
		AccountID:     account.ID,
		TransactionID: tx.ID,
		CreditDebit:   "DR",
		Amount:        tx.Amount,
		BalanceAfter:  account.Balance - tx.Amount,
	}

	entry, err := testQueries.CreateLedgerEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.TransactionID, entry.TransactionID)
	require.Equal(t, arg.CreditDebit, entry.CreditDebit)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.BalanceAfter, entry.BalanceAfter)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}