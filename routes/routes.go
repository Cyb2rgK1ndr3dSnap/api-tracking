package routes

import (
	"net/http"

	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/controllers"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/security"
	"github.com/gin-gonic/gin"
)

func RootRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Tracking is running!",
		})
	})
}

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("", security.AuthMiddleware())
		userGroup.POST("", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
	}
}

func ShippingRoutes(r *gin.Engine) {
	shippingGroup := r.Group("/shipping")
	{
		shippingGroup.GET("", security.AuthMiddleware(), controllers.GetShipping)
		shippingGroup.POST("", security.AuthMiddleware(), security.AdminMiddleware(), controllers.CreateShipping)
		shippingGroup.PUT("", security.AuthMiddleware(), security.AdminMiddleware(), controllers.UpdateShipping)
		shippingGroup.POST("/close", security.AuthMiddleware(), security.AdminMiddleware(), controllers.CloseShipping)
	}
}

func TransactionRoutes(r *gin.Engine) {
	transactionGroup := r.Group("/transaction")
	{
		transactionGroup.GET("", security.AuthMiddleware(), controllers.BalanceTransaction)
		transactionGroup.POST("", security.AuthMiddleware(), security.AdminMiddleware(), controllers.CreateTransaction)
	}
}
