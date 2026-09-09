package db

import (
	"context"
	"testing"
	"time"

	"github.com/afsar-hussai/apex_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		FullName:     util.RandomFullName(),
		Email:        util.RandomEmail(),
		Phone:        util.RandomPhone(),
		PasswordHash: "hashed_password_xyz",
		Role:         util.RandomRole(),
		KycStatus:    util.RandomKycStatus(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.KycStatus, user.KycStatus)
	require.True(t, user.IsActive)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Phone, user2.Phone)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}