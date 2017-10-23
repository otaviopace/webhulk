package postback

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
		mutex: sync.RWMUtex{},
	}, nil
}

func (s *PostbackStore) Store(p *Postback) (*Postback, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, err := db.NamedExec(`
		INSERT INTO postback
		(id, status, payload, created_at)
		VALUES (:id, :status, :payload, :created_at)
	`, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PostbackStore) Retrieve(id string) (*Postback, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	p := &Postback{}

	err := db.Get(p, `SELECT * FROM postbacks WHERE id = $1`, id)

	if err != nil {
		return nil, false
	}

	return p
}
