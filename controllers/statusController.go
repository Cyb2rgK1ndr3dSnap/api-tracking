package controllers

import (
	"database/sql"
	"time"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	cache := models.NewSimpleCache()

	// Define una lista de 10 colores predefinidos
	predefinedColors := []string{
		"#D4EDDA", "#F8D7DA", "#CDD5FF", "#FFF3CD", "#000000",
	}

	// Verifica si los estados están en la caché
	if cachedStatuses, found := cache.Get("statuses"); found {
		c.JSON(200, gin.H{"statuses": cachedStatuses})
		return
	}

	// Obtén los estados de la "base de datos"
	statuses, err := services.GetStatus(db)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with database"})
		return
	}

	// Asigna un color a cada estado
	for i := range statuses {
		colorIndex := i % len(predefinedColors) // Índice cíclico para los colores
		statuses[i].Color = predefinedColors[colorIndex]
	}

	// Guarda en la caché con un TTL de 60 minutos
	cache.Set("statuses", statuses, 60*time.Minute)

	// Envía los estados con sus colores
	c.JSON(200, statuses)
}
