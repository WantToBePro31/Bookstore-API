package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func CreateDB() *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		panic("Could not load .env file")
	}
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open("mysql", connectionInfo)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})

	return db
}
