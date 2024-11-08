package routes

import (
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/controllers"
	"github.com/gin-gonic/gin"
)

func ShippingRoutes(r *gin.Engine) {
	shippingGroup := r.Group("/shipping")
	{
		shippingGroup.POST("", controllers.CreateShipping)
	}
}

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("", controllers.CreateUser)

	}
}
