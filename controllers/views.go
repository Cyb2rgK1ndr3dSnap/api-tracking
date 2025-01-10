package controllers

import (
	"github.com/gin-gonic/gin"
)

func ViewLogin(c *gin.Context) {

	c.HTML(200, "login.html", nil)
}

func ViewIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ViewTracking(c *gin.Context) {
	c.HTML(200, "track.html", nil)
}
