package api

import (
	"github.com/Owoade/go-bank/config"
	"github.com/Owoade/go-bank/service"
	"github.com/Owoade/go-bank/sql"
	"github.com/Owoade/go-bank/token"
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
			Store:      store,
			ConfigVars: configVars,
		},
	}

	tokenMaker, _ := token.NewPasetomaker(configVars.PasetoSymetricToken)

	authRoute := server.router.Group("/").Use(authMidleware(tokenMaker))

	server.router.POST("/auth/login", server.login)

	server.router.POST("/auth/signup", server.signup)

	authRoute.POST("/bank/account", server.createAccount)

	authRoute.POST("/bank/account/credit", server.creditAccount)

	authRoute.POST("/bank/account/transfer", server.transferCash)

	return server
}

func (server *Server) Start(address string) error {

	return server.router.Run(address)

}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
