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
	FromAccountId int64    `json:"from_account" binding:"required,min=1"`
	ToAccountId   int64    `json:"to_account" binding:"required,min=1"`
	Amount        *big.Int `json:"amount" binding:"required"`
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

func (s *Service) CreateAccount(ctx context.Context, userId int32) (sql.Account, error) {

	createUserParams := sql.CreateAccountParams{
		UserID: pgtype.Int4{
			Int32: userId,
			Valid: true,
		},
		Balance: pgtype.Numeric{
			Int:   big.NewInt(0),
			Valid: true,
		},
	}

	account, err := s.Store.Queries.CreateAccount(ctx, createUserParams)

	return account, err

}

func (s *Service) CreditAccount(ctx context.Context, amount *big.Int, accountId int64) (sql.Account, error) {

	arg := sql.CreditAccountParams{
		Balance: pgtype.Numeric{
			Int:   amount,
			Valid: true,
		},
		ID: accountId,
	}

	account, err := s.Store.Queries.CreditAccount(ctx, arg)

	return account, err

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

	fmt.Println(sender.Balance.Int, arg.Amount)

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
