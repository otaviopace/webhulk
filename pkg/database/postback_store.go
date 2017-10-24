package database

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

type PostbackStore struct {
	db    *sqlx.DB
	mutex sync.RWMutex
}

func NewPostbackStore(db *sqlx.DB) (*PostbackStore, error) {
	return &PostbackStore{
		db:    db,
		mutex: sync.RWMutex{},
	}, nil
}

func (s *PostbackStore) Store(p *Postback) (*Postback, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	query := `
		INSERT INTO postbacks
			(id, status, payload, created_at)
		VALUES
			(:id, :status, :payload, :created_at)
	`

	_, err := s.db.NamedExec(query, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PostbackStore) Retrieve(id string) (*Postback, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	p := &Postback{}

	query := `
		SELECT
			id, status, payload, created_at
		FROM postbacks
		WHERE
			id = $1
	`

	err := s.db.Get(p, query, id)

	if err != nil {
		return nil, false
	}

	return p, true
}
