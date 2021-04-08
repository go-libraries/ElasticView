package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

func runEsLink(app *gin.Engine) {
	esLink := app.Group("/api/es_link")
	{
		esLink.POST("InsertAction", EsLinkController{}.InsertAction)
		esLink.POST("DeleteAction", EsLinkController{}.DeleteAction)
		esLink.POST("UpdateAction", EsLinkController{}.UpdateAction)
		esLink.GET("ListAction", EsLinkController{}.ListAction)
	}
}
