package database

import (
	"goapi-project-structure/pkg/config"
	"goapi-project-structure/pkg/module/todos/core/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func New(conf *config.Config) (*gorm.DB, error) {
	// Build a DSN e.g. postgres://username:password@host:port/dbName
	// dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v",
	// 	conf.Db.Driver,
	// 	conf.Db.Username,
	// 	conf.Db.Password,
	// 	conf.Db.Host,
	// 	conf.Db.Port,
	// 	conf.Db.Database,
	// 	conf.Db.Sslmode)

	gcnf := &gorm.Config{}

	if conf.App.Mode == "production" {
		gcnf.Logger = logger.Default.LogMode(logger.Silent)
	} else {
		gcnf.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(sqlite.Open(conf.Db.Database), gcnf)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(model.Todo{})
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
