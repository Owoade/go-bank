package seeders

import (
	"context"
	"log"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

func (store *SeederStore) AccountNumber(accountId int32) sql.AccountNumber {

	payload := sql.CreateAccountNumberParams{
		AccountID: pgtype.Int4{
			Int32: accountId,
			Valid: true,
		},
		AccountName: pgtype.Text{
			String: utils.GenerateRandomString(12),
			Valid:  true,
		},
		AccountNumber: pgtype.Int8{
			Int64: utils.GenerateRandomInteger(5),
			Valid: true,
		},
		BankName: pgtype.Text{
			String: utils.GenerateRandomString(9),
			Valid:  true,
		},
	}

	accountNumber, err := store.queries.CreateAccountNumber(context.Background(), payload)

	if err != nil {
		log.Fatal("error seeding db with account numbers")
	}

	return accountNumber

}

func (store *SeederStore) AccountNumbers(accountId int32, n int) []sql.AccountNumber {

	accountNumbers := make([]sql.AccountNumber, 0, n)

	for i := 1; i <= n; i++ {

		accountNumber := store.AccountNumber(accountId)

		accountNumbers = append(accountNumbers, accountNumber)

	}

	return accountNumbers

}
