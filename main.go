package main

import (
	"flag"
	"go-cms/config"
	"go-cms/routes"
	"go-cms/seeder"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	if handleFlags() {
		return
	}
	startServer()
}

func handleFlags() bool {
	db := config.ConnectDatabase()
	seed := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()
	if *seed {
		log.Println("Running database seeder...")
		seeder.SeedDatabase(db)
		return true
	}
	return false
}

func startServer() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	port := config.GetPort()
	log.Printf("Starting server on port %s...", port)
	r.Run(":" + port)
}
