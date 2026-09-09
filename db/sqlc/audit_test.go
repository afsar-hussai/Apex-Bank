package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateAuditLog(t *testing.T) {
	user := createRandomUser(t)

	arg := CreateAuditLogParams{
		ActorID:   user.ID,
		Action:    "UPDATE",
		TableName: "users",
		IpAddress: pgtype.Text{String: "192.168.1.1", Valid: true},
	}

	logEntry, err := testQueries.CreateAuditLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, logEntry)

	require.Equal(t, arg.Action, logEntry.Action)
	require.Equal(t, arg.TableName, logEntry.TableName)
	require.NotZero(t, logEntry.ID)
	require.NotZero(t, logEntry.CreatedAt)
}