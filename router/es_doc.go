package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES文档 路由
func runEsDoc(app *gin.Engine) {
	esIndex := app.Group("/api/es_doc")
	{
		esIndex.POST("DeleteRowByIDAction", EsDocController{}.DeleteRowByIDAction)
		esIndex.POST("UpdateByIDAction", EsDocController{}.UpdateByIDAction)
		esIndex.POST("InsertAction", EsDocController{}.InsertAction)

	}
}
