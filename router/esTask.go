package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

func runEsTask(app *gin.Engine) {
	task := app.Group("/api/es_task")
	{
		task.POST("ListAction", TaskController{}.ListAction)
		task.POST("CancelAction", TaskController{}.CancelAction)
	}
}
