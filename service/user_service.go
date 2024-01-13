package service

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/Owoade/go-bank/sql"
	"github.com/jackc/pgx/v5/pgtype"
)

type TransferCashParams struct {
	ToAccountId   int64
	Amount        *big.Int
	FromAccountId int64
}

type TransferCashResult struct {
	CreditTransaction sql.Transaction
	DebitTransaction  sql.Transaction
	DebitedAccount    sql.Account
	CreditedAcccount  sql.Account
	Recipient         sql.Account
	Sender            sql.Account
}

type TransferCashTransactionCallback struct {
	tc_params TransferCashParams
	ctx       context.Context
	Recipient sql.Account
	Sender    sql.Account
}

func (params TransferCashTransactionCallback) transferCashDbTransaction(q *sql.Queries) (interface{}, error) {
	arg := params.tc_params
	ctx := params.ctx

	sender := params.Sender
	recipient := params.Recipient

	amountValue := pgtype.Numeric{
		Int:   arg.Amount,
		Valid: true,
	}

	debitAccountParams := sql.DebitAccountParams{
		ID:      arg.FromAccountId,
		Balance: amountValue,
	}

	creditAcountParams := sql.CreditAccountParams{
		ID:      arg.ToAccountId,
		Balance: amountValue,
	}

	creditTransactionParams := sql.CreateTransactionParams{
		AccountID: pgtype.Int4{
			Int32: int32(arg.ToAccountId),
			Valid: true,
		},
		Amount: amountValue,
		Type: sql.NullTransactionStatus{
			TransactionStatus: "credit",
			Valid:             true,
		},
	}

	debitTransactionParams := sql.CreateTransactionParams{
		AccountID: pgtype.Int4{
			Int32: int32(arg.FromAccountId),
			Valid: true,
		},
		Amount: amountValue,
		Type: sql.NullTransactionStatus{
			TransactionStatus: "debit",
			Valid:             true,
		},
	}

	debitedAccount, debitAccountError := q.DebitAccount(ctx, debitAccountParams)

	creditedAccount, creditAccountError := q.CreditAccount(ctx, creditAcountParams)

	creditTransaction, creditTransactionError := q.CreateTransaction(ctx, creditTransactionParams)

	debitTransaction, debitTransactionError := q.CreateTransaction(ctx, debitTransactionParams)

	if debitAccountError != nil {
		log.Fatal("falied to debit account")
	}

	if creditAccountError != nil {
		log.Fatal("failed to credit account")
	}

	if creditTransactionError != nil {
		log.Fatal("failed to create credit transaction")
	}

	if debitTransactionError != nil {
		log.Fatal("failed to create debit transaction")
	}

	result := TransferCashResult{
		DebitTransaction:  debitTransaction,
		CreditTransaction: creditTransaction,
		DebitedAccount:    debitedAccount,
		CreditedAcccount:  creditedAccount,
		Recipient:         recipient,
		Sender:            sender,
	}

	fmt.Println(result.CreditTransaction.CreatedAt)

	return result, nil

}

func (s *Service) TransferCash(ctx context.Context, arg TransferCashParams) (TransferCashResult, error) {

	zeroValue := *new(TransferCashResult)

	sender, get_sender_err := s.Store.Queries.GetAccountById(ctx, arg.FromAccountId)

	recipient, get_recipient_err := s.Store.Queries.GetAccountById(ctx, arg.ToAccountId)

	if get_sender_err != nil {
		return zeroValue, fmt.Errorf("faliled to get sender account")
	}

	if get_recipient_err != nil {
		return zeroValue, fmt.Errorf("faliled to get sender account")
	}

	expectedBalanceAfterTransaction := sender.Balance.Int.Cmp(arg.Amount)

	if expectedBalanceAfterTransaction < 0 {
		return zeroValue, fmt.Errorf("insuffucient funds")
	}

	transactionCallback := TransferCashTransactionCallback{
		tc_params: arg,
		ctx:       ctx,
		Recipient: recipient,
		Sender:    sender,
	}

	result, err := s.Store.ExecTx(ctx, transactionCallback.transferCashDbTransaction)

	if err != nil {
		return zeroValue, fmt.Errorf("failed to execute transaction: %v", err)
	}

	// Check if the result is nil before type assertion
	if result == nil {
		return zeroValue, fmt.Errorf("transaction result is nil")
	}

	// Type assertion with error handling
	typeCastedResult, ok := result.(TransferCashResult)
	if !ok {
		return zeroValue, fmt.Errorf("unexpected result type")
	}

	return typeCastedResult, nil
}
