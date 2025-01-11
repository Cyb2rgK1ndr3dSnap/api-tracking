package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewLogin(c *gin.Context) {

	_, err := c.Cookie("jwt_token")

	if err != nil {
		c.HTML(200, "login.html", nil)
		c.Abort()
		return
	}

	c.Redirect(http.StatusFound, "/index")
}

func ViewIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ViewTracking(c *gin.Context) {
	c.HTML(200, "track.html", nil)
}
