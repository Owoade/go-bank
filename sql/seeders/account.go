package seeders

import (
	"context"
	"log"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

func (store *SeederStore) Account(userId int32) sql.Account {

	arg := sql.CreateAccountParams{
		UserID: pgtype.Int4{
			Int32: userId,
			Valid: true,
		},
		Balance: pgtype.Numeric{
			Int:   utils.GenerateRandomInteger(3),
			Valid: true,
		},
	}

	account, err := store.queries.CreateAccount(context.Background(), arg)

	if err != nil {
		log.Fatal("error seeding db with accounts")
	}

	return account

}
