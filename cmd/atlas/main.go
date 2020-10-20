package main

import (
	"github.com/soganhei.com.br/atlas/handler"
	"github.com/soganhei.com.br/atlas/jwt"
	"github.com/soganhei.com.br/atlas/postgres"
)

func main() {

	jwtServices, err := jwt.NewServices()
	if err != nil {
		panic(err)
	}

	services := postgres.NewServices()

	server := handler.Server{
		JwtServices:   jwtServices,
		UsersServices: services.UsersServices,
	}
	server.Start()

}
