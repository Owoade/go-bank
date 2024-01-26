package api

import (
	"github.com/Owoade/go-bank/config"
	"github.com/Owoade/go-bank/service"
	"github.com/Owoade/go-bank/sql"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	service    *service.Service
	configVars *config.Variables
}

func NewServer(store *sql.SQLStore) *Server {

	configVars := config.NewVars()

	server := &Server{
		router: gin.Default(),
		service: &service.Service{
			Store: store,
		},
		configVars: configVars,
	}

	server.router.POST("/auth/login", server.login)

	server.router.POST("/auth/signup", server.signup)

	server.router.POST("/bank/account", server.createAccount)

	server.router.POST("/bank/account/credit", server.creditAccount)

	server.router.POST("/bank/account/transfer", server.transferCash)

	return server
}

func (server *Server) Start(address string) error {

	return server.router.Run(address)

}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
