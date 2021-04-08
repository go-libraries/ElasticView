package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

func runGmGuid(app *gin.Engine) {
	guid := app.Group("/api/gm_guid")
	{
		guid.POST("Finish", GuidController{}.Finish)
		guid.POST("IsFinish", GuidController{}.IsFinish)
	}
}
