package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
// InitDB инициализирует подключение к базе данных
func InitDB() {
    dsn := "host=localhost user=postgres password=password dbname=postgres port=5433 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
}
