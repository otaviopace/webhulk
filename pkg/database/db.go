package database

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	cfg "github.com/grvcoelho/webhulk/pkg/config"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase(conf *cfg.Database) (*Database, error) {
	db, err := sqlx.Open("postgres", conf.Address)

	if err != nil {
		log.WithFields(log.Fields{
			"address": conf.Address,
			"error":   err,
		}).Error()

		return nil, err
	}

	return &Database{db}, nil
}
