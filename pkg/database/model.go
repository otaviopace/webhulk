package database

import (
	"time"

	cuid "gopkg.in/lucsky/cuid.v1"
)

type Model struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

func NewId() string {
	return cuid.New()
}

func NewIdWithPrefix(prefix string) string {
	return prefix + "_" + NewId()
}
