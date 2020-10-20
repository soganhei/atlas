package postgres

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/soganhei.com.br/atlas/postgres/users"

	_ "github.com/lib/pq"
)

type Services struct {
	db            *sqlx.DB
	UsersServices *users.Services
}

func NewServices() *Services {

	services := Services{}

	dataSourceName := os.Getenv("DATABASE_PGSQL_URL")

	if dataSourceName == "" {
		panic("DATABASE_PGSQL_URL is empty")
	}

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	users := users.NewServices(db)
	services.UsersServices = users

	return &services
}
