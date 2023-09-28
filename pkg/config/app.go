package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBUserName = "your_username"
	DBPassword = "your_password"
	DBName     = "your_db_name"
)

func Connect() *gorm.DB {
	dsn := DBUserName + ":" + DBPassword + "@tcp(localhost:3306)/" + DBName + "?parseTime=True"

	// Open a connection to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}
