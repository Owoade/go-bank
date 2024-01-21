package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Owoade/go-bank/api"
	"github.com/Owoade/go-bank/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	dbSource := os.Getenv("DB_URL")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	fmt.Println(dbSource, serverAddress)

	connPool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("Couldn't connect wirh database")
	}

	store := sql.NewStore(connPool)

	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Couldn't start server")
	}

}
