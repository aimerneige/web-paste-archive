package route

import (
	"github.com/aimerneige/web-paste/controller"
	"github.com/gin-gonic/gin"
)

func AppRouteCollection(r *gin.Engine) *gin.Engine {
	r.LoadHTMLFiles("app/index.html")
	r.StaticFS("/static", gin.Dir("app/static", true))

	r.GET("/", controller.AppIndex)

	return r
}
