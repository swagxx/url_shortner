package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"judo/configs"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *DB {
	db := OpenDB(conf)
	return &DB{
		db,
	}

}

func OpenDB(conf *configs.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
