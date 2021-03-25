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
		guid.Any("Finish", GuidController{}.Finish)
		guid.Any("IsFinish", GuidController{}.IsFinish)
	}

	esLink := app.Group("/api/es_link")
	{
		esLink.Any("InsertAction", EsLinkController{}.InsertAction)
		esLink.Any("DeleteAction", EsLinkController{}.DeleteAction)
		esLink.Any("UpdateAction", EsLinkController{}.UpdateAction)
		esLink.Any("ListAction", EsLinkController{}.ListAction)
	}

	es := app.Group("/api/es")
	{
		es.Any("PingAction", EsController{}.PingAction)
		es.Any("CatAction", EsController{}.CatAction)
		es.Any("RunDslAction", EsController{}.RunDslAction)
		es.Any("SqlToDslAction", EsController{}.SqlToDslAction)
	}
	esMap := app.Group("/api/es_map")
	{
		esMap.Any("ListAction", EsMappingController{}.ListAction)
	}

	esIndex := app.Group("/api/es_index")
	{
		esIndex.Any("CreateAction", EsIndexController{}.CreateAction)
		esIndex.Any("GetSettingsAction", EsIndexController{}.GetSettingsAction)
		esIndex.Any("IndexNamesAction", EsIndexController{}.IndexNamesAction)
		esIndex.Any("ReindexAction", EsIndexController{}.ReindexAction)
		esIndex.Any("GetAliasAction", EsIndexController{}.GetAliasAction)
		esIndex.Any("OperateAliasAction", EsIndexController{}.OperateAliasAction)
	}

	dslHistory := app.Group("/api/dslHistory")
	{
		dslHistory.Any("CleanAction", DslHistoryController{}.CleanAction)
		dslHistory.Any("ListAction", DslHistoryController{}.ListAction)
	}

	return app
}
