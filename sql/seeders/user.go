package seeders

import (
	"context"
	"fmt"
	"log"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type SeededUser struct {
	sql.User
	RawPassword string
}

func (store *SeederStore) User() SeededUser {

	randomPassword := utils.GenerateRandomString(10)

	hashedPassword, err := utils.HashPassword(randomPassword)

	if err != nil {
		log.Fatal("Passsword hashing failed:", err)
	}

	arg := sql.CreateUserParams{
		Email: pgtype.Text{
			String: fmt.Sprintf("%s%s", utils.GenerateRandomString(8), "@go-bank.com"),
			Valid:  true,
		},
		Password: pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		},
	}

	user, err := store.queries.CreateUser(context.Background(), arg)

	if err != nil {
		log.Fatal("error seeding db with user")
	}

	return SeededUser{
		User:        user,
		RawPassword: randomPassword,
	}

}
