package controllers

import (
	"database/sql"
	"time"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/models"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/services"
	"github.com/gin-gonic/gin"
)

func GetLocker(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	cache := models.NewSimpleCache()

	// Verifica si los estados están en la caché
	if cachedLocker, found := cache.Get("locker"); found {
		c.JSON(200, cachedLocker)
		return
	}

	// Obtén los estados de la "base de datos"
	locker, err := services.GetLocker(db)
	if err != nil {
		c.JSON(400, gin.H{"error": "server error"})
		return
	}

	// Guarda en la caché con un TTL de 60 minutos
	cache.Set("locker", locker, 60*time.Minute)

	// Envía los estados con sus colores
	c.JSON(200, locker)
}
