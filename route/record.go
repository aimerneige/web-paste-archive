package route

import (
	"github.com/aimerneige/web-paste/controller"
	"github.com/aimerneige/web-paste/middleware"
	"github.com/gin-gonic/gin"
)

func RecordRouteCollection(g *gin.RouterGroup) {
	g.POST("/",
		middleware.RecordCreateRequestCheck(),
		controller.CreateNewRecord)

	g.GET("/:id",
		middleware.RecordPathIDCheck(),
		controller.GetRecordByID)

	g.DELETE("/:id",
		middleware.RecordPathIDCheck(),
		controller.DeleteRecordByID)
}
