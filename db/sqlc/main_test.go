package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	
	// Local test database connection string
	connString := "postgres://root:secret@localhost:5432/apex_bank?sslmode=disable"
	
	testDB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	err = testDB.Ping(context.Background())
	if err != nil {
		log.Fatal("cannot ping db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}