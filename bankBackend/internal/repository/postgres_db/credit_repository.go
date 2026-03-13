package core

import "database/sql"

type CreditRepository struct {
	db *sql.DB
}