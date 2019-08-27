package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var PG *gorm.DB

func init() {
	PG = connectDB()
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=dev password=123456 dbname=local_0722 port=5432 sslmode=disable")
	if err != nil {
		log.Error("Database connect error:", err)
	}

	db.SingularTable(true)

	return db
}
