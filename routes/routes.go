package routes

import (
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/controllers"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/security"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
	}
}

func ShippingRoutes(r *gin.Engine) {
	shippingGroup := r.Group("/shipping")
	{
		shippingGroup.GET("", security.AuthMiddleware(), controllers.ReadShipping)
		shippingGroup.POST("", security.AuthMiddleware(), controllers.CreateShipping)
		shippingGroup.PUT("", security.AuthMiddleware(), controllers.UpdateShipping)
	}
}
