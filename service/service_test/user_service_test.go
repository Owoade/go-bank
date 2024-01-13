package service_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/Owoade/go-bank/service"
	"github.com/Owoade/go-bank/sql"
	"github.com/stretchr/testify/require"
)

func TestTransferCash(t *testing.T) {

	arg := service.TransferCashParams{
		ToAccountId:   1,
		FromAccountId: 3,
		Amount:        big.NewInt(10),
	}

	result, err := testService.TransferCash(context.Background(), arg)

	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, err)

	require.Equal(t, result.CreditTransaction.Type.TransactionStatus, sql.TransactionStatus("credit"))
	require.Equal(t, result.CreditTransaction.AccountID.Int32, int32(arg.ToAccountId))

	require.Equal(t, result.DebitTransaction.Type.TransactionStatus, sql.TransactionStatus("debit"))
	require.Equal(t, result.DebitTransaction.AccountID.Int32, int32(arg.FromAccountId))

	fmt.Println(result.CreditTransaction.Amount.Int, arg.Amount, result.CreditTransaction.AccountID.Int32)
	require.Equal(t, result.CreditTransaction.Amount.Int.Cmp(arg.Amount), 0)
	require.Equal(t, result.DebitTransaction.Amount.Int.Cmp(arg.Amount), 0)

	require.Equal(t, new(big.Int).Sub(result.Sender.Balance.Int, arg.Amount).Cmp(result.DebitedAccount.Balance.Int), 0)
	require.Equal(t, new(big.Int).Add(result.Recipient.Balance.Int, arg.Amount).Cmp(result.CreditedAcccount.Balance.Int), 0)

}
