package sql_test

import (
	"context"
	"log"
	"math/big"
	"testing"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	ctx := context.Background()

	user := sqlSeeders.User()

	arg := sql.CreateAccountParams{
		UserID: pgtype.Int4{
			Int32: user.ID,
			Valid: true,
		},
		Balance: pgtype.Numeric{
			Int:   utils.GenerateRandomInteger(3),
			Valid: true,
		},
	}

	account, err := testQueries.CreateAccount(ctx, arg)

	require.NoError(t, err)
	require.Equal(t, arg.Balance, account.Balance)

	require.Equal(t, account.UserID, arg.UserID)
	require.NotEmpty(t, account.CreatedAt)

}

func TestGetUserAccount(t *testing.T) {

	userId := pgtype.Int4{
		Int32: 1,
		Valid: true,
	}

	account, err := testQueries.GetUserAccount(context.Background(), userId)

	require.NoError(t, err)
	require.Equal(t, account.UserID, userId)

}

func TestCreditAccount(t *testing.T) {

	randomAmount := utils.GenerateRandomInteger(3)

	user := sqlSeeders.User()

	account := sqlSeeders.Account(user.ID)

	payload := sql.CreditAccountParams{
		ID: account.ID,
		Balance: pgtype.Numeric{
			Int:   randomAmount,
			Valid: true,
		},
	}

	updatedAccount, err := testQueries.CreditAccount(context.Background(), payload)

	if err != nil {
		log.Fatal("update account query failed")
	}

	updatedAccountBalance := updatedAccount.Balance.Int

	existingAccountBalance := account.Balance.Int

	balanceDifference := new(big.Int).Sub(updatedAccountBalance, existingAccountBalance)

	require.NoError(t, err)

	require.Equal(t, randomAmount, balanceDifference)

	require.Equal(t, account.ID, updatedAccount.ID)

}
