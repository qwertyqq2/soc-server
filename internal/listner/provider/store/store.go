package store

import (
	"database/sql"

	"github.com/qwertyqq2/soc-server/internal/listner/provider/model"
)

type Store struct {
	db          *sql.DB
	rep         *Repository
	PlayersChan chan *model.Player
	LotsChan    chan *model.Lot
}

func New(db *sql.DB) *Store {
	pchan := make(chan *model.Player, 10)
	lchan := make(chan *model.Lot, 10)
	return &Store{
		db:          db,
		PlayersChan: pchan,
		LotsChan:    lchan,
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
