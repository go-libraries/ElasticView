package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES 任务 路由
func runEsTask(app *gin.Engine) {
	task := app.Group("/api/es_task")
	{
		task.POST("ListAction", TaskController{}.ListAction)
		task.POST("CancelAction", TaskController{}.CancelAction)
	}
}
