package store

import (
	"github.com/go-pg/pg/v9"
)

type StoreConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func NewStore(config *pg.Options) *pg.DB {
	db := pg.Connect(config)

	return db
}
