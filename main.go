package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Owoade/go-bank/api"
	"github.com/Owoade/go-bank/config"
	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/token"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	dbSource := os.Getenv("DB_URL")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	fmt.Println(dbSource, serverAddress)

	fmt.Println(config.NewVars().PasetoTokenDuration)

	maker, err := token.NewPasetomaker(utils.GenerateRandomString(32))

	token, _ := maker.CreateToken(1, config.NewVars().PasetoTokenDuration)

	fmt.Println(token)

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
