package token

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey []byte
}

const minSkLen = 32

func NewJWTMaker(sk string) (*JWTMaker, error) {

	fmt.Println(sk)

	if len(sk) < minSkLen {
		return nil, fmt.Errorf("secret key must be at least %d characters", minSkLen)
	}

	return &JWTMaker{secretKey: []byte(sk)}, nil

}

func (maker *JWTMaker) CreateToken(id int64, duration time.Duration) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	payload, _err := NewPayload(id, duration)

	fmt.Println(_err)

	claims["exp"] = time.Minute

	claims["user"] = payload.UserId

	claims["id"] = payload.Id

	claims["iat"] = payload.ExpiresAt

	tokenString, err := token.SignedString(maker.secretKey)

	fmt.Println(err)

	return tokenString, err

}

func (maker *JWTMaker) VerifyToken(tokenString string) (Payload, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return maker.secretKey, nil
	})

	if err != nil {
		return *new(Payload), err
	}

	if !token.Valid {
		return *new(Payload), fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		user := claims["user"].(float64)
		iat := claims["iat"].(string)
		id := claims["id"].(string)
		exp := claims["exp"].(float64)

		tokenPayload := Payload{
			UserId:    int64(user),
			Id:        id,
			ExpiresAt: strconv.FormatFloat(exp, 'f', -1, 64),
			IssuedAt:  iat,
		}

		return tokenPayload, nil

	}

	return *new(Payload), fmt.Errorf("unable to extract payload from token")

}
