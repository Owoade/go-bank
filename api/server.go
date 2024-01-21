package api

import (
	"github.com/Owoade/go-bank/service"
	"github.com/Owoade/go-bank/sql"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	service *service.Service
}

func NewServer(store *sql.SQLStore) *Server {

	server := &Server{
		router: gin.Default(),
		service: &service.Service{
			Store: store,
		},
	}

	server.router.POST("/auth/login", server.login)

	server.router.POST("/auth/signup", server.signup)

	return server
}

func (server *Server) Start(address string) error {

	return server.router.Run(address)

}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}