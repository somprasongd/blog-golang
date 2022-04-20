package database

import (
	"fmt"
	"goapi/pkg/common/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	db := config.Config.Db
	// Build a DSN e.g. postgres://username:password@host:port/dbName
	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v",
		db.Driver,
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Database)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}

	log.Println("DB Connected")
}
