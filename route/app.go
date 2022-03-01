package route

import (
	"github.com/aimerneige/web-paste/controller"
	"github.com/gin-gonic/gin"
)

func AppRouteCollection(r *gin.Engine) *gin.Engine {
	r.LoadHTMLFiles("app/index.html", "app/access.html")
	r.StaticFS("/static", gin.Dir("app/static", true))

	r.GET("/", controller.AppIndex)
	r.GET("/:link", controller.AppAccess)

	return r
}
