package router

import (
	. "github.com/1340691923/ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES连接 路由
func runEsLink(app *gin.Engine) {
	esLink := app.Group("/api/es_link")
	{
		esLink.POST("InsertAction", EsLinkController{}.InsertAction)
		esLink.POST("DeleteAction", EsLinkController{}.DeleteAction)
		esLink.POST("UpdateAction", EsLinkController{}.UpdateAction)
		esLink.GET("ListAction", EsLinkController{}.ListAction)
	}
}
