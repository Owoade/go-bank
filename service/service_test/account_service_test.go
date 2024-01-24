package service_test

import (
	"context"
	"log"
	"testing"

	"github.com/Owoade/go-bank/service"
	"github.com/Owoade/go-bank/sql"
	"github.com/stretchr/testify/require"
)

func TestTransferCash(t *testing.T) {

	sender := sqlSeeders.User()

	recipient := sqlSeeders.User()

	fromAccount := sqlSeeders.Account(sender.ID)

	toAccount := sqlSeeders.Account(recipient.ID)

	arg := service.TransferCashParams{
		ToAccountId:   toAccount.ID,
		FromAccountId: fromAccount.ID,
		Amount:        20,
	}

	result, err := testService.TransferCash(context.Background(), arg)

	if err != nil {
		log.Fatal(err)
	}

	creditTransactionAmount := result.CreditTransaction.Amount.Int64

	debitTransactionAmount := result.DebitTransaction.Amount.Int64

	require.NoError(t, err)

	require.Equal(t, result.CreditTransaction.Type.TransactionStatus, sql.TransactionStatus("credit"))
	require.Equal(t, result.CreditTransaction.AccountID.Int32, int32(arg.ToAccountId))

	require.Equal(t, result.DebitTransaction.Type.TransactionStatus, sql.TransactionStatus("debit"))
	require.Equal(t, result.DebitTransaction.AccountID.Int32, int32(arg.FromAccountId))

	require.Equal(t, creditTransactionAmount, arg.Amount)
	require.Equal(t, debitTransactionAmount, arg.Amount)

	require.Equal(t, (result.Sender.Balance.Int64 - arg.Amount), result.DebitedAccount.Balance.Int64)
	require.Equal(t, (result.Recipient.Balance.Int64 + arg.Amount), result.CreditedAcccount.Balance.Int64)

}
