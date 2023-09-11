package config

import (
	"example/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(dsn string) *gorm.DB {
	// dsn := "host=localhost user=postgres password=postgres dbname=p2-w3 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.AutoMigrate(&model.Player{}, &model.User{}); err != nil {
		log.Fatal(err.Error())
	}

	return db
}
