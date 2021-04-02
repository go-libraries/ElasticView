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
	gmUser := app.Group("/api/gm_user")
	{
		gmUser.Any("info", UserController{}.UserInfo)
		gmUser.Any("roles", RoleController{}.RolesAction)
		gmUser.Any("role/update", RoleController{}.RolesUpdateAction)
		gmUser.Any("role/add", RoleController{}.RolesAddAction)
		gmUser.Any("role/delete", RoleController{}.RolesDelAction)
		gmUser.Any("logout", UserController{}.LogoutAction)
		gmUser.Any("userlist", UserController{}.UserListAction)
		gmUser.Any("roleOption", RoleController{}.RoleOptionAction)
		gmUser.Any("getUserById", UserController{}.GetUserByIdAction)
		gmUser.Any("UpdateUser", UserController{}.UserUpdateAction)
		gmUser.Any("InsertUser", UserController{}.UserAddAction)
		gmUser.Any("DelUser", UserController{}.DeleteUserAction)
	}

	guid := app.Group("/api/gm_guid")
	{
		guid.POST("Finish", GuidController{}.Finish)
		guid.POST("IsFinish", GuidController{}.IsFinish)
	}

	esLink := app.Group("/api/es_link")
	{
		esLink.POST("InsertAction", EsLinkController{}.InsertAction)
		esLink.POST("DeleteAction", EsLinkController{}.DeleteAction)
		esLink.POST("UpdateAction", EsLinkController{}.UpdateAction)
		esLink.GET("ListAction", EsLinkController{}.ListAction)
	}

	es := app.Group("/api/es")
	{
		es.POST("PingAction", EsController{}.PingAction)
		es.POST("CatAction", EsController{}.CatAction)
		es.POST("RunDslAction", EsController{}.RunDslAction)
		es.Any("SqlToDslAction", EsController{}.SqlToDslAction)
	}
	esMap := app.Group("/api/es_map")
	{
		esMap.POST("ListAction", EsMappingController{}.ListAction)
	}

	esIndex := app.Group("/api/es_index")
	{
		esIndex.POST("CreateAction", EsIndexController{}.CreateAction)
		esIndex.POST("GetSettingsAction", EsIndexController{}.GetSettingsAction)
		esIndex.POST("IndexNamesAction", EsIndexController{}.IndexNamesAction)
		esIndex.POST("ReindexAction", EsIndexController{}.ReindexAction)
		esIndex.POST("GetAliasAction", EsIndexController{}.GetAliasAction)
		esIndex.POST("OperateAliasAction", EsIndexController{}.OperateAliasAction)
	}

	dslHistory := app.Group("/api/dslHistory")
	{
		dslHistory.POST("CleanAction", DslHistoryController{}.CleanAction)
		dslHistory.POST("ListAction", DslHistoryController{}.ListAction)
	}

	task := app.Group("/api/Task")
	{
		task.POST("ListAction", TaskController{}.ListAction)
	}

	backUp := app.Group("/api/backUp")
	{
		backUp.POST("SnapshotListAction", EsBackUpController{}.SnapshotListAction)
	}

	return app
}
