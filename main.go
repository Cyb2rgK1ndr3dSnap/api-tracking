package main

import (
	"log"
	"os"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/initializers"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")

	db := initializers.InitDB()
	defer db.Close()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	go routes.UserRoutes(r)
	routes.ShippingRoutes(r)

	r.Run(":" + port)
}
