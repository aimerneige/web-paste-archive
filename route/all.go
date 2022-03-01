package route

import (
	"github.com/aimerneige/web-paste/controller"
	"github.com/aimerneige/web-paste/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r.SetTrustedProxies([]string{
		"127.0.0.1",
	})
	r.Use(middleware.BlackListMiddleware())

	record := r.Group("/record")
	RecordRouteCollection(record)

	staticPath := viper.GetString("common.path")
	r.Static("/raw", staticPath)

	AppRouteCollection(r)

	r.NoRoute(controller.NotFound)

	return r
}
