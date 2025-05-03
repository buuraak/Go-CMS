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
	seed := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()

	config.LoadEnv()
	db := config.ConnectDatabase()

	if *seed {
		log.Println("Running database seeder...")
		seeder.SeedDatabase(db)
		return
	}

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":" + config.GetPort())
}
