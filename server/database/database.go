package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	dialect := mysql.Open("root:@tcp(127.0.0.1:3306)/first_web_with_golang_playground?charset=utf8mb4&parseTime=True&loc=Local")
	DB, err = gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to Database...")
	}

	fmt.Println("Connecting to database...")
}