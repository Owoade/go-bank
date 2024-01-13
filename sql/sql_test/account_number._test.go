package sql_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountnumber(t *testing.T) {

	accountId := pgtype.Int4{
		Int32: 1,
		Valid: true,
	}

	accountNumbers, err := generateRandomAccountNumbers(t, 4, accountId)

	if err != nil {
		log.Fatal("error creating account numbers")
	}

	require.NoError(t, err)
	require.Len(t, accountNumbers, 4)

	for i := 0; i < len(accountNumbers); i++ {

		accountNumber := accountNumbers[i]

		require.NotEmpty(t, accountNumber)

	}

}

func TestGetUserAccountNumbers(t *testing.T) {

	accountId := pgtype.Int4{
		Int32: 1,
		Valid: true,
	}

	accountNumbers, err := generateRandomAccountNumbers(t, 5, accountId)

	if err != nil {
		log.Fatal("error creating account numbers")
	}

	require.NoError(t, err)
	require.Len(t, accountNumbers, 5)

}

func generateRandomAccountNumbers(t *testing.T, n int, accountId pgtype.Int4) ([]sql.AccountNumber, error) {

	accountNumbers := make([]sql.AccountNumber, 0, n)

	for i := 1; i <= n; i++ {

		payload := sql.CreateAccountNumberParams{
			AccountID: pgtype.Int4{
				Int32: 1,
				Valid: true,
			},
			AccountName: pgtype.Text{
				String: utils.GenerateRandomString(12),
				Valid:  true,
			},
			AccountNumber: pgtype.Int8{
				Int64: utils.GenerateRandomInteger(5).Int64(),
				Valid: true,
			},
			BankName: pgtype.Text{
				String: utils.GenerateRandomString(9),
				Valid:  true,
			},
		}

		accountNumber, err := testQueries.CreateAccountNumber(context.Background(), payload)

		if err != nil {
			return accountNumbers, fmt.Errorf("error creating account number %s", fmt.Sprint(i))
		}

		accountNumbers = append(accountNumbers, accountNumber)

	}

	return accountNumbers, nil

}
