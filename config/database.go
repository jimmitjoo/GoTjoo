package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UseDatabase = true

func ConnectDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/gotjoo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	return db
}
