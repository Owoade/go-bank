package service_test

import (
	"fmt"
	"testing"

	"github.com/Owoade/go-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {

	existingUser := sqlSeeders.User()

	successResponse, successErr := testService.Login(existingUser.Email.String, existingUser.RawPassword)

	_, err := testService.Login(utils.GenerateRandomString(5), utils.GenerateRandomString(9))

	require.Error(t, err)
	require.NoError(t, successErr)

	require.NotEmpty(t, successResponse)

}

func TestSignUp(t *testing.T) {

	existingUser := sqlSeeders.User()

	randomEmail := fmt.Sprintf("%s%s", utils.GenerateRandomString(8), "@go-bank.com")

	randomPassword := utils.GenerateRandomString(10)

	successResponse, successErr := testService.SignUp(randomEmail, randomPassword)

	_, err := testService.SignUp(existingUser.Email.String, existingUser.Password.String)

	require.Error(t, err)
	require.NoError(t, successErr)

	require.Equal(t, successResponse.Email.String, randomEmail)
	require.True(t, utils.CompareHashedPassword(successResponse.Password.String, randomPassword))

}
