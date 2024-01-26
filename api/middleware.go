package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Owoade/go-bank/token"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey  = "authorization"
	authHeaderType = "bearer"
)

func authMidleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader(authHeaderKey)

		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse(fmt.Errorf("authorization header not found")))
			return
		}

		fields := strings.Fields(authHeader)

		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse(fmt.Errorf("invalid auth header format")))
			return
		}

		authHeaderTypeBearer := strings.ToLower(fields[0])

		if authHeaderType != authHeaderTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse(fmt.Errorf("invalid auth header type")))
			return
		}

		token := fields[1]
		payload, err := tokenMaker.VerifyToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse(fmt.Errorf("invalid token")))
			return
		}

		ctx.Set(authHeaderKey, payload)

	}

}
