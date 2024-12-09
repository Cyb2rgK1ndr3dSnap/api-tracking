package security

import (
	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		roleID, _ := c.Get("roleID")

		if roleID != 1 {
			c.JSON(401, gin.H{"error": "Unauthorized, not admin"})
			c.Abort()
			return
		}
		c.Next()
	}
}
