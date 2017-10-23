package database

import (
	"sync"
	"time"

	cuid "gopkg.in/lucsky/cuid.v1"
)

type Postback struct {
	Model
	Status  string `db:"status"`
	Payload string `db:"payload"`
}

func NewPostback(status, payload string) *Postback {
	return &Postback{
		ID:        cuid.New(),
		Status:    status,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}

type PostbackStore struct {
	cache map[string]*Postback
	mutex sync.RWMutex
}

func NewPostbackStore() (*PostbackStore, error) {
	return &PostbackStore{
		cache: make(map[string]*Postback, 0),
		mutex: sync.RWMutex{},
	}, nil
}

func (s *PostbackStore) Store(p *Postback) (*Postback, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache[p.ID] = p

	return p, nil
}

func (s *PostbackStore) Retrieve(id string) (*Postback, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	p, ok := s.cache[id]

	if !ok {
		return nil, false
	}

	return p, true
}

func (s *PostbackStore) Delete(id string) (bool, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.cache, id)
	return true, nil
}
