package core

import "database/sql"

type CurrencyRepository struct {
	db *sql.DB
}