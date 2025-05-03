package seeder

import (
	"go-cms/models"
	"log"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User model: %s", err)
	}

	log.Println("Database seeding completed successfully!")
}
