package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AppAccess(c *gin.Context) {
	c.HTML(http.StatusOK, "access.html", nil)
}
