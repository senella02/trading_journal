package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=admin password=root dbname=DooDuang port=5433 sslmode=disable TimeZone=Asia/Bangkok"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
}
