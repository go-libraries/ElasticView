package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES索引 路由
func runEsIndex(app *gin.Engine) {
	esIndex := app.Group("/api/es_index")
	{
		esIndex.POST("DeleteAction", EsIndexController{}.DeleteAction)
		esIndex.POST("CreateAction", EsIndexController{}.CreateAction)
		esIndex.POST("GetSettingsAction", EsIndexController{}.GetSettingsAction)
		esIndex.POST("IndexNamesAction", EsIndexController{}.IndexNamesAction)
		esIndex.POST("ReindexAction", EsIndexController{}.ReindexAction)
		esIndex.POST("GetAliasAction", EsIndexController{}.GetAliasAction)
		esIndex.POST("OperateAliasAction", EsIndexController{}.OperateAliasAction)
		esIndex.POST("GetSettingsInfoAction", EsIndexController{}.GetSettingsInfoAction)
		esIndex.POST("StatsAction", EsIndexController{}.StatsAction)

	}
}
