package db

import (
	"context"
	"testing"
	"time"

	"github.com/afsar-hussai/apex_bank/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T) Transaction {
	senderAccount := createRandomAccount(t)
	receiverAccount := createRandomAccount(t)

	arg := CreateTransactionParams{
		SenderAccountID:   pgtype.Int8{Int64: senderAccount.ID, Valid: true},
		ReceiverAccountID: pgtype.Int8{Int64: receiverAccount.ID, Valid: true},
		Amount:            util.RandomMoney(),
		Currency:          senderAccount.Currency,
		Status:            "completed",
		TransactionType:   util.RandomTransactionType(),
		IdempotencyKey:    pgtype.Text{String: util.RandomString(16), Valid: true},
		Description:       pgtype.Text{String: "Test fund transfer", Valid: true},
	}

	tx, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tx)

	require.Equal(t, arg.SenderAccountID, tx.SenderAccountID)
	require.Equal(t, arg.ReceiverAccountID, tx.ReceiverAccountID)
	require.Equal(t, arg.Amount, tx.Amount)
	require.Equal(t, arg.Currency, tx.Currency)
	require.Equal(t, arg.Status, tx.Status)
	require.Equal(t, arg.TransactionType, tx.TransactionType)

	require.NotZero(t, tx.ID)
	require.NotZero(t, tx.CreatedAt)

	return tx
}

func TestCreateTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestGetTransactionById(t *testing.T) {
	tx1 := createRandomTransaction(t)
	tx2, err := testQueries.GetTransactionById(context.Background(), tx1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, tx2)

	require.Equal(t, tx1.ID, tx2.ID)
	require.Equal(t, tx1.Amount, tx2.Amount)
	require.Equal(t, tx1.Status, tx2.Status)
	require.WithinDuration(t, tx1.CreatedAt.Time, tx2.CreatedAt.Time, time.Second)
}