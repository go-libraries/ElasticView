package router

import (
	. "github.com/1340691923/ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// DslHistory 路由
func runDslHistory(app *gin.Engine) {
	dslHistory := app.Group("/api/dslHistory")
	{
		dslHistory.POST("CleanAction", DslHistoryController{}.CleanAction)
		dslHistory.POST("ListAction", DslHistoryController{}.ListAction)
	}

}
