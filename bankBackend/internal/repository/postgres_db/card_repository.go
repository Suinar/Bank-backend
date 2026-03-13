package core

import "database/sql"

type CardRepository struct {
	db *sql.DB
}
