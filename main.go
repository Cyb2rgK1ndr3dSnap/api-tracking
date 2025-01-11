package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/Cyb2rgK1ndr3dSnap/api-tracking/docs"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/initializers"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/routes"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/tasks"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed templates/*
var htmlFiles embed.FS

//go:embed static/*
var staticFiles embed.FS

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

	// Cargar plantillas desde el binario
	tmpl := template.Must(template.New("").ParseFS(htmlFiles, "templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	// Servir archivos estáticos desde el binario
	staticFileServer := http.FileServer(http.FS(staticFiles))

	r.GET("/static/*filepath", func(c *gin.Context) {
		// Remover el prefijo /static para que coincida con la estructura del sistema embed
		filepath := c.Param("filepath")
		c.Request.URL.Path = "static" + filepath // Agregar solo el prefijo "static"
		staticFileServer.ServeHTTP(c.Writer, c.Request)
	})

	// Servir favicon desde el sistema embed
	r.GET("/favicon.ico", func(c *gin.Context) {
		favicon, err := staticFiles.ReadFile("static/favicon.ico")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "image/x-icon", favicon)
	})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	go tasks.StartScheduler(db)

	routes.ViewRoutes(r)
	routes.UserRoutes(r)
	routes.ShippingRoutes(r)
	routes.TransactionRoutes(r)
	routes.BusinessRoutes(r)
	routes.LockerRoutes(r)
	routes.StatusRoutes(r)

	log.Println("Server is running on port:", port)

	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
