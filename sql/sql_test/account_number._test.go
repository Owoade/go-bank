package sql_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccountnumber(t *testing.T) {

	user := sqlSeeders.User()

	account := sqlSeeders.Account(user.ID)

	accountNumbers := sqlSeeders.AccountNumbers(int32(account.ID), 5)

	require.Len(t, accountNumbers, 5)

	for i := 0; i < len(accountNumbers); i++ {

		accountNumber := accountNumbers[i]

		require.NotEmpty(t, accountNumber)

		require.Equal(t, accountNumber.AccountID.Int32, int32(account.ID))

	}

}
