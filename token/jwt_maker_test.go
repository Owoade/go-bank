package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/Owoade/go-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {

	sk := []byte(utils.GenerateRandomString(32))

	fmt.Println(sk)

	user := int64(1)

	jwtMaker := JWTMaker{
		secretKey: sk,
	}

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := jwtMaker.CreateToken(user, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := jwtMaker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, payload.UserId, user)

	// Define the layout that matches the datetime string
	layout := "2006-01-02 15:04:05.9999999 -0700 MST m=-1.020304"
	_issuedAt, _ := time.Parse(layout, issuedAt.String())
	__issuedAt, _ := time.Parse(layout, payload.IssuedAt)

	_expiredAt, _ := time.Parse(layout, expiredAt.String())
	__expireAt, _ := time.Parse(layout, expiredAt.String())

	require.WithinDuration(t, _issuedAt, __issuedAt, time.Second)
	require.WithinDuration(t, _expiredAt, __expireAt, time.Second)

}
