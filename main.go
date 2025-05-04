package main

import (
	"flag"
	"go-cms/config"
	"go-cms/routes"
	"go-cms/seeder"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDatabase()
	if handleFlags(db) {
		return
	}
	startServer(db)
}

func handleFlags(db *gorm.DB) bool {
	seed := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()
	if *seed {
		log.Println("Running database seeder...")
		seeder.SeedDatabase(db)
		return true
	}
	return false
}

func startServer(db *gorm.DB) {
	r := gin.Default()
	routes.RegisterRoutes(r, db)
	port := config.GetPort()
	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
