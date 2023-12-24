// Package sqlc
// Created by GoLand.
// User: nixon
// Date: 24/12/2023
// Time: 05:12
package sqlc

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
	*Queries
}

func (s *Store) NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
