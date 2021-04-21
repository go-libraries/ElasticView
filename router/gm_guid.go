package router

import (
	. "github.com/1340691923/ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES 新手引导 路由
func runGmGuid(app *gin.Engine) {
	guid := app.Group("/api/gm_guid")
	{
		guid.POST("Finish", GuidController{}.Finish)
		guid.POST("IsFinish", GuidController{}.IsFinish)
	}
}
