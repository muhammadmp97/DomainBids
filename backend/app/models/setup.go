package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	MySQL "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	config := MySQL.Config{
		User:   os.Getenv("DATABAE_USERNAME"),
		Passwd: os.Getenv("DATABAE_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DATABASE_HOST"),
		DBName: os.Getenv("DATABASE_NAME"),
	}

	database, err := gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Auction{})

	DB = database
}
