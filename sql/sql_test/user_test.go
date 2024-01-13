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

func TestCreateUser(t *testing.T) {

	hashedPassword, err := utils.HashPassword("owoade anu")

	if err != nil {
		log.Fatal("Passsword hashing failed:", err)
	}

	arg := sql.CreateUserParams{
		Email: pgtype.Text{
			String: fmt.Sprintf("%s%s", utils.GenerateRandomString(8), "@go-bank.com"),
			Valid:  true,
		},
		Password: pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, user.Password, arg.Password)

	require.NotZero(t, user.ID)

}
