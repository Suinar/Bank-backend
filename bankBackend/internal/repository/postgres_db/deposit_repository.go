package core

import "database/sql"

type DepositRepository struct {
	db *sql.DB
}