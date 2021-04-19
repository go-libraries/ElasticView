package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES基础操作 路由
func runEs(app *gin.Engine) {
	es := app.Group("/api/es")
	{
		es.POST("RecoverCanWrite", EsController{}.RecoverCanWrite)
		es.POST("PingAction", EsController{}.PingAction)
		es.POST("CatAction", EsController{}.CatAction)
		es.POST("RunDslAction", EsController{}.RunDslAction)
		es.Any("SqlToDslAction", EsController{}.SqlToDslAction)
		es.POST("OptimizeAction", EsController{}.OptimizeAction)
	}
}
