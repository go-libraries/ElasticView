package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

// ES备份 路由
func runEsBackUp(app *gin.Engine) {
	backUp := app.Group("/api/backUp")
	{
		backUp.POST("SnapshotRepositoryListAction", EsBackUpController{}.SnapshotRepositoryListAction)
		backUp.POST("SnapshotListAction", EsBackUpController{}.SnapshotListAction)
		backUp.POST("SnapshotCreateRepositoryAction", EsBackUpController{}.SnapshotCreateRepositoryAction)
		backUp.POST("SnapshotDeleteRepositoryAction", EsBackUpController{}.SnapshotDeleteRepositoryAction)
		backUp.POST("CleanupeRepositoryAction", EsBackUpController{}.CleanupeRepositoryAction)
		backUp.POST("CreateSnapshotAction", EsBackUpController{}.CreateSnapshotAction)
		backUp.POST("SnapshotDeleteAction", EsBackUpController{}.SnapshotDeleteAction)
		backUp.POST("SnapshotDetailAction", EsBackUpController{}.SnapshotDetailAction)
		backUp.POST("SnapshotRestoreAction", EsBackUpController{}.SnapshotRestoreAction)
		backUp.POST("SnapshotStatusAction", EsBackUpController{}.SnapshotStatusAction)

	}
}
