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

	if *seed {
		log.Println("Running database seeder...")
		seeder.SeedDatabase()
		return
	}

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":" + config.GetEnv("PORT", "8080"))
}
