package store

import "database/sql"

type sqlStore struct {
	db *sql.DB
}

