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
	runEsBackUp(app)
	runEsDoc(app)
	return app
}
