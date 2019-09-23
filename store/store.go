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
	// dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	config.Host, config.Port, config.UserName, config.Password, config.DbName)
	// db, err := sql.Open("postgres", dbinfo)

	// if err != nil {
	// 	return nil
	// }

	// return db
	db := pg.Connect(config)

	return db
}
