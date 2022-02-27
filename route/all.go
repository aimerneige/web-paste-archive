package route

import (
	"github.com/aimerneige/web-paste/controller"
	"github.com/gin-gonic/gin"
)

func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r.NoRoute(controller.NotFound)

	return r
}
