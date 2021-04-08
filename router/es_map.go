package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

func runEsMap(app *gin.Engine) {
	esMap := app.Group("/api/es_map")
	{
		esMap.POST("ListAction", EsMappingController{}.ListAction)
	}
}
