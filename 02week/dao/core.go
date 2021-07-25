package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	globalDB *gorm.DB

	globalConfig *DBConfig
)

func Connect(cfg *DBConfig) {
	globalConfig = cfg

	db, err := gorm.Open(sqlite.Open(globalConfig.DBPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// // Migrate the schema
	// db.AutoMigrate(&model.User{})

	globalDB = db
}

func GetDB() (db *gorm.DB) {
	return globalDB
}
