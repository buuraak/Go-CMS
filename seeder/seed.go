package seeder

import (
	"go-cms/config"
	"go-cms/models"
	"log"
)

func SeedDatabase() {
	db := config.ConnectDatabase()

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User model: %s", err)
	}

	log.Println("Database seeding completed successfully!")
}
