package core

import "database/sql"

type AccountRepository struct {
	db *sql.DB
}
