package database

import (
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/config"
)

type Database struct {
	*gorm.DB
}

func New(config *config.Config) (*Database, error) {
	dsn := "user=" + config.GetString("DB_USERNAME") + " password=" + config.GetString("DB_PASSWORD") + " dbname=" + config.GetString("DB_DATABASE") + " host=" + config.GetString("DB_HOST") + " port=" + strconv.Itoa(config.GetInt("DB_PORT")) + " TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			tz, _ := time.LoadLocation("Asia/Jakarta")

			return time.Now().In(tz)
		},
	})

	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
