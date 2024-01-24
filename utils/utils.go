package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	math_rand "math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	math_rand.Seed(time.Now().UnixNano())
}

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

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func generateRandomNumber(min, max int) int {
	return math_rand.Intn(max-min+1) + min
}

func GenerateRandomInteger(length int) int64 {

	if length <= 0 {
		return *new(int64)
	}

	// Seed the random number generator with the current time
	min := intPow(10, length-1)
	max := intPow(10, length) - 1
	return int64(generateRandomNumber(min, max))

}

func GetRandomValueFromSLice[T any](s []T) T {

	if len(s) == 0 {
		return *new(T)
	}

	index := math_rand.Intn(len(s))

	return s[index]

}

func CompareHashedPassword(pass1, pass2 string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(pass1), []byte(pass2))

	if err == nil {
		return true
	}

	fmt.Println(err.Error())

	return false

}
