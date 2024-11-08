package main

import (
	"log"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/initializers"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := initializers.InitDB()
	defer db.Close()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.ShippingRoutes(r)

	r.Run(":8000")
}
