package database

import (
	configs "github.com/Suinar/Bank-backend/bankBackend/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDb(cfg *configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.Postgres.DbUrl)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)

	return db, nil
}
