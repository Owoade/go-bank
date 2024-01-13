package sql

import (
	"log"
	"os"
	"testing"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/bank-app?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	connPool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("Couldn't connect wirh database")
	}

	testQueries = New(connPool)

	os.Exit(m.Run())
}
