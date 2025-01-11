package routes

import (
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/controllers"
	"github.com/Cyb2rgK1ndr3dSnap/api-tracking/security"
	"github.com/gin-gonic/gin"
)

func RootRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Indaboxs is alive!",
		})
	})
}

func ViewRoutes(r *gin.Engine) {
	//Views
	r.GET("/", controllers.ViewLogin)
	r.GET("/index", security.AuthMiddleware(), controllers.ViewIndex)
	r.GET("/track", security.AuthMiddleware(), controllers.ViewTracking)
}

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/test", security.AuthMiddleware())
		userGroup.POST("", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.POST("/logout", security.AuthMiddleware(), controllers.LogoutUser)
		//Web API
		userGroup.GET("", security.AuthMiddleware(), security.AdminMiddleware(), controllers.GetUsers)
		userGroup.POST("/total", security.AuthMiddleware(), security.AdminMiddleware(), controllers.GetUsersTotal)
	}
}

func ShippingRoutes(r *gin.Engine) {
	shippingGroup := r.Group("/shipping")
	{
		shippingGroup.GET("", security.AuthMiddleware(), controllers.GetShipping)
		shippingGroup.POST("", security.AuthMiddleware(), security.AdminMiddleware(), controllers.CreateShipping)
		shippingGroup.PUT("", security.AuthMiddleware(), security.AdminMiddleware(), controllers.UpdateShipping)
		shippingGroup.POST("/close", security.AuthMiddleware(), security.AdminMiddleware(), controllers.CloseShipping)
		//Web API
		shippingGroup.POST("/total", security.AuthMiddleware(), security.AdminMiddleware(), controllers.GetShippingsTotal)
	}
}

func TransactionRoutes(r *gin.Engine) {
	transactionGroup := r.Group("/transaction")
	{
		transactionGroup.GET("", security.AuthMiddleware(), controllers.BalanceTransaction)
		transactionGroup.POST("", security.AuthMiddleware(), security.AdminMiddleware(), controllers.CreateTransaction)
	}
}

func BusinessRoutes(r *gin.Engine) {
	businessGroup := r.Group("/business")
	{
		businessGroup.GET("", security.AuthMiddleware(), controllers.GetBusinessInformation)
	}
}

func LockerRoutes(r *gin.Engine) {
	lockerGroup := r.Group("/locker")
	{
		lockerGroup.GET("", security.AuthMiddleware(), controllers.GetLocker)
	}
}

func StatusRoutes(r *gin.Engine) {
	statusGroup := r.Group("/status")
	{
		statusGroup.GET("", security.AuthMiddleware(), controllers.GetStatus)
	}
}
