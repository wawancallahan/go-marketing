package database

import (
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/config"
)

type Database struct {
	*gorm.DB
}

func New(config *config.Config) *Database {
	dsn := "user=" + config.GetString("DB_USERNAME") + " password=" + config.GetString("DB_PASSWORD") + " dbname=" + config.GetString("DB_DATABASE") + " host=" + config.GetString("DB_HOST") + " port=" + strconv.Itoa(config.GetInt("DB_PORT")) + " TimeZone=UTC"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &Database{db}
}
