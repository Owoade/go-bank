package sql

import (
	"database/sql"
	main_sql "database/sql"
	"github/Owoade/go-bank/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/bank-app?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := main_sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Couldn't connect wirh database")
	}

	testQueries = sql.New(conn)

	os.Exit(m.Run())
}
