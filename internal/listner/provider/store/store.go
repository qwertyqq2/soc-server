package store

import (
	"database/sql"
)

type Store struct {
	db  *sql.DB
	rep *Repository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Repository() *Repository {
	if s.rep != nil {
		return s.rep
	}
	s.rep = &Repository{
		store: s,
	}

	return s.rep
}

func (s *Store) Close() {
	s.db.Close()
}
