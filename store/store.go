package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type StoreConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func NewStore(config StoreConfig) *sql.DB {

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.UserName, config.Password, config.DbName)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil
	}

	return db
}
