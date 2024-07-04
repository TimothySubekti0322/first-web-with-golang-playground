package migrations

import (
	"fmt"
	"log"

	"github.com/TimothySubekti0322/first-web-with-golang-playground/database"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/models/entity"
)

func Migration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal("Failed to Migrate Database...")
	}

	fmt.Print("Migrated Successfully")
}