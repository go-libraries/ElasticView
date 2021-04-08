package router

import (
	. "ElasticView/controller"
	"ElasticView/platform-basic-libs/util"
	_ "ElasticView/statik"

	. "ElasticView/middleware"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

// statik -src=views/dist -f
func Init() *gin.Engine {
	app := gin.Default()
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	app.GET("/", func(context *gin.Context) {
		htmlStr := `
			<html>
				<body>	
					<a href="view" style="display: block;text-align: center;">Hello ElasticView ! GO -></a>
				</body>
			</html>
		`
		context.Writer.Write(util.Str2bytes(htmlStr))
	})

	app.StaticFS("/view", statikFS)

	app.Use(Cors)
	app.Any("/api/gm_user/login", UserController{}.Login)
	app.Use(JwtMiddleware)
	runGmUser(app)
	runGmGuid(app)
	runEsLink(app)
	runEs(app)
	runEsMap(app)
	runEsIndex(app)
	runDslHistory(app)
	runEsTask(app)

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

	return app
}
