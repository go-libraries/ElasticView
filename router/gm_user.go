package router

import (
	. "ElasticView/controller"

	"github.com/gin-gonic/gin"
)

func runGmUser(app *gin.Engine) {
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
}
