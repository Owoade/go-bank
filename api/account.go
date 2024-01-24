package api

import (
	"net/http"

	"github.com/Owoade/go-bank/service"
	"github.com/gin-gonic/gin"
)

type CreateAccountParams struct {
	UserId int32 `json:"userId" binding:"required"`
}

type CreditAccountParams struct {
	AccountId int64 `json:"accountId" binding:"required"`
	Amount    int64 `json:"amount" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {

	var req CreateAccountParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	response, err := server.service.CreateAccount(ctx, req.UserId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	ctx.JSON(http.StatusOK, response)

}

func (server *Server) transferCash(ctx *gin.Context) {

	var req service.TransferCashParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	response, err := server.service.TransferCash(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	ctx.JSON(http.StatusOK, response)

}

func (server *Server) creditAccount(ctx *gin.Context) {

	var req CreditAccountParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	response, err := server.service.CreditAccount(ctx, req.Amount, req.AccountId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	ctx.JSON(http.StatusOK, response)

}
