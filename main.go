package main

import (
	"log"
	"os"

	_ "github.com/Cyb2rgK1ndr3dSnap/api-tracking/docs"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/initializers"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Tag Service API
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

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

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	// add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.RootRoute(r)
	routes.UserRoutes(r)
	routes.ShippingRoutes(r)
	routes.TransactionRoutes(r)

	log.Println("Server is running on port:", port)

	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
