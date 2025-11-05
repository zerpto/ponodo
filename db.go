package core

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormConnection(host, port, user, password, dbName string) *gorm.DB {
	if port == "" {
		port = "5432"
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
