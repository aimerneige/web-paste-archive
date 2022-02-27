package controller

import (
	"github.com/aimerneige/web-paste/response"
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	response.NotFound(c, nil, "404 page not found")
}
