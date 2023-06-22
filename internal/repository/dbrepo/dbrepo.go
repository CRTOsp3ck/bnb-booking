package dbrepo

import (
	"bnb-booking/internal/config"
	"bnb-booking/internal/repository"
	"database/sql"
)

type postgresDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDbRepo{
		App: a,
	}
}
