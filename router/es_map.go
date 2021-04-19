package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES mapping 路由
func runEsMap(app *gin.Engine) {
	esMap := app.Group("/api/es_map")
	{
		esMap.POST("ListAction", EsMappingController{}.ListAction)
		esMap.POST("UpdateMappingAction", EsMappingController{}.UpdateMappingAction)

	}
}
