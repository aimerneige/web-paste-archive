package middleware

import (
	"github.com/aimerneige/web-paste/dto"
	"github.com/aimerneige/web-paste/response"
	"github.com/gin-gonic/gin"
)

// RecordCreateRequestCheck check request body
func RecordCreateRequestCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.RecordRequestDto
		err := c.ShouldBindJSON(&request)
		if err != nil {
			response.BadRequest(c, err, "Fail to bind request body.")
			c.Abort()
			return
		}

		if request.Content == "" {
			response.BadRequest(c, nil, "Content is required.")
			c.Abort()
			return
		}

		if request.Title == "" {
			request.Title = "Untitled"
		}

		c.Set("content", request.Content)
		c.Set("title", request.Title)
		c.Next()
	}
}

// RecordPathIDCheck check path id
func RecordPathIDCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			response.BadRequest(c, nil, "ID is required.")
			c.Abort()
			return
		}

		c.Set("id", id)
		c.Next()
	}
}

// RecordPathLinkCheck check path link
func RecordPathLinkCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		link := c.Param("link")
		if link == "" {
			response.BadRequest(c, nil, "Link is required.")
			c.Abort()
			return
		}

		c.Set("link", link)
		c.Next()
	}
}
