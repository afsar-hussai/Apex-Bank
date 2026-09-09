package db

import (
	"context"
	"testing"
	"time"

	"github.com/afsar-hussai/apex_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		UserID:      user.ID,
		Balance:     util.RandomMoney(),
		Currency:    util.RandomCurrency(),
		AccountType: util.RandomAccountType(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.AccountType, account.AccountType)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.AccountType, account2.AccountType)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestAddAccountBalance(t *testing.T) {
	account1 := createRandomAccount(t)
	anAmount := int64(500)

	updatedAccount, err := testQueries.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		Amount:    anAmount,
		AccountID: account1.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, account1.Balance+anAmount, updatedAccount.Balance)
}