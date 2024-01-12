package utils

import (
	"crypto/rand"
	"math/big"
	math_rand "math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		result += string(charset[randomIndex.Int64()])
	}

	return result
}

func GenerateRandomInteger(length int) *big.Int {

	if length <= 0 {
		return *new(*big.Int)
	}

	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	randomInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil
	}

	return randomInt

}

func GetRandomValueFromSLice[T any](s []T) T {

	if len(s) == 0 {
		return *new(T)
	}

	index := math_rand.Intn(len(s))

	return s[index]

}
