package sql_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	arg := sql.CreateAccountParams{
		UserID: pgtype.Int4{
			Int32: 2,
			Valid: true,
		},
		Balance: pgtype.Numeric{
			Int:   utils.GenerateRandomInteger(3),
			Valid: true,
		},
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

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

	payload := sql.CreditAccountParams{
		ID: 1,
		Balance: pgtype.Numeric{
			Int:   randomAmount,
			Valid: true,
		},
	}

	existingAccount, e_err := testQueries.GetAccountById(context.Background(), 1)

	updatedAccount, err := testQueries.CreditAccount(context.Background(), payload)

	fmt.Println(existingAccount.Balance.Int, updatedAccount.Balance.Int)
	if e_err != nil {
		fmt.Println(e_err)
		log.Fatal("get account query failed")
	}

	if err != nil {
		log.Fatal("update account query failed")
	}

	updatedAccountBalance := updatedAccount.Balance.Int

	existingAccountBalance := existingAccount.Balance.Int

	balanceDifference := new(big.Int).Sub(updatedAccountBalance, existingAccountBalance)

	require.NoError(t, err)

	require.Equal(t, randomAmount, balanceDifference)

}
