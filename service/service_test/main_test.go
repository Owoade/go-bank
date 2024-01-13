package service_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"context"

	"github.com/Owoade/go-bank/service"
	"github.com/Owoade/go-bank/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/bank-app?sslmode=disable"
)

var testQueries *sql.Queries

var testService *service.Service

func TestMain(m *testing.M) {

	connPool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("Couldn't connect wirh database")
	}

	testService = &service.Service{
		Store: sql.NewStore(connPool),
	}

	fmt.Println(&testService.Store, "test service")

	testQueries = sql.New(connPool)

	os.Exit(m.Run())
}
