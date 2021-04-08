package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

func runEsTask(app *gin.Engine) {
	task := app.Group("/api/Task")
	{
		task.POST("ListAction", TaskController{}.ListAction)
	}
}
