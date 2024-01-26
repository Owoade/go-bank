package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/token"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type LoginResponse struct {
	id    int32
	token string
}

func (s *Service) Login(email string, password string) (LoginResponse, error) {

	typeCastedEmail := pgtype.Text{
		String: email,
		Valid:  true,
	}

	existingUser, err := s.Store.Queries.GetOneUserByEmail(context.Background(), typeCastedEmail)

	if err != nil {
		return *new(LoginResponse), fmt.Errorf(err.Error())
	}

	if !utils.CompareHashedPassword(existingUser.Password.String, password) {
		return *new(LoginResponse), fmt.Errorf("incorect password")
	}

	pacetoMaker, perr := token.NewPasetomaker(s.ConfigVars.PasetoSymetricToken)

	if perr != nil {
		return *new(LoginResponse), fmt.Errorf("error creating token")
	}

	token, err := pacetoMaker.CreateToken(int64(existingUser.ID), s.ConfigVars.PasetoTokenDuration.Abs())

	if err != nil {
		return *new(LoginResponse), fmt.Errorf("error creating token")
	}

	response := LoginResponse{
		id:    existingUser.ID,
		token: token,
	}

	return response, nil

}

func (s *Service) SignUp(email string, password string) (sql.User, error) {

	typeCastedEmail := pgtype.Text{
		String: email,
		Valid:  true,
	}

	hashedPassword, _ := utils.HashPassword(password)

	typeCastedPassword := pgtype.Text{
		String: hashedPassword,
		Valid:  true,
	}

	existingUser, err := s.Store.Queries.GetOneUserByEmail(context.Background(), typeCastedEmail)

	if err == nil {
		return *new(sql.User), errors.New(fmt.Sprintf("user with email %s already exists", existingUser.Email.String))
	}

	createUserParams := sql.CreateUserParams{
		Email:    typeCastedEmail,
		Password: typeCastedPassword,
	}

	newUser, newUserErr := s.Store.Queries.CreateUser(context.Background(), createUserParams)

	if newUserErr != nil {
		return *new(sql.User), errors.New(fmt.Sprintf("acouldn;t create new user: %s", err.Error()))
	}

	return newUser, nil

}
