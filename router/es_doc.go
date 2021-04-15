package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

func runEsDoc(app *gin.Engine) {
	esIndex := app.Group("/api/es_doc")
	{
		esIndex.POST("DeleteRowByIDAction", EsDocController{}.DeleteRowByIDAction)
	}
}
