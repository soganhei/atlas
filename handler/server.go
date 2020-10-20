package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/soganhei.com.br/atlas"
	"github.com/soganhei.com.br/atlas/handler/auth"
	"github.com/soganhei.com.br/atlas/handler/users"
)

type Server struct {
	JwtServices   atlas.JwtServices
	UsersServices atlas.UsersServices
}

func (server *Server) Start() {

	router := gin.Default()

	r := router.Group("/")

	jwt := server.JwtServices

	r.Use(authTokenJwt(jwt))
	{

		users := users.NewHandler(r)
		users.UsersServices = server.UsersServices
	}

	auth := auth.NewHandler(r)
	auth.UsersServices = server.UsersServices
	auth.JwtServices = server.JwtServices

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
