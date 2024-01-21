package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (server *Server) login(ctx *gin.Context) {

	var req UserAuthParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	response, err := server.service.Login(req.Email, req.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	ctx.JSON(http.StatusOK, response)

}

func (server *Server) signup(ctx *gin.Context) {

	var req UserAuthParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	response, err := server.service.SignUp(req.Email, req.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	ctx.JSON(http.StatusOK, response)

}
