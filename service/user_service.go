package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) Login(email string, password string) (int32, error) {

	typeCastedEmail := pgtype.Text{
		String: email,
		Valid:  true,
	}

	existingUser, err := s.Store.Queries.GetOneUserByEmail(context.Background(), typeCastedEmail)

	if err != nil {
		return *new(int32), fmt.Errorf(err.Error())
	}

	if !utils.CompareHashedPassword(existingUser.Password.String, password) {
		return *new(int32), fmt.Errorf("incorect password")
	}

	return existingUser.ID, nil

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
