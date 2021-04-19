package controller

import (
	"ElasticView/engine/logs"
	"ElasticView/model"
	"ElasticView/platform-basic-libs/jwt"
	"ElasticView/platform-basic-libs/my_error"
	"ElasticView/platform-basic-libs/response"
	"ElasticView/platform-basic-libs/service/gm_user"

	"github.com/gin-gonic/gin"
)

//GM用户控制器
type UserController struct {
	BaseController
}

// 登录
func (this UserController) Login(ctx *gin.Context) {
	username := ctx.Request.FormValue("username")
	password := ctx.Request.FormValue("password")
	var gmUserService gm_user.GmUserService
	token, err := gmUserService.CheckLogin(username, password)
	if err != nil {
		logs.Logger.Sugar().Errorf("登陆失败", err)
		err = my_error.NewBusiness(gm_user.AUTH_ERROR, gm_user.ERROR_AUTH)
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "登陆成功", map[string]interface{}{"token": token})
}

// 用户详细信息
func (this UserController) UserInfo(ctx *gin.Context) {
	var gmUserService gm_user.GmUserService

	token := this.GetToken(ctx)
	claims, err := jwt.ParseToken(token)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	info, err := gmUserService.GetRoleInfo(claims.RoleId)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "登陆成功", map[string]interface{}{"roles": []string{"admin"}, "introduction": info.Description, "name": info.RoleName, "list": info.RoleList, "avatar": ""})
}

//退出登录
func (this UserController) LogoutAction(ctx *gin.Context) {
	this.Success(ctx, response.LogoutSuccess, nil)
}

//GM 用户列表
func (this UserController) UserListAction(ctx *gin.Context) {
	var userModel model.GmUserModel
	list, err := userModel.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}

// 删除GM用户
func (this UserController) DeleteUserAction(ctx *gin.Context) {
	var userModel model.GmUserModel
	userModel.ID = int32(this.FormIntDefault(ctx, "id", 0))
	err := userModel.Delete()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.DeleteSuccess, nil)
}

// 用ID获取用户信息
func (this UserController) GetUserByIdAction(ctx *gin.Context) {
	var userModel model.GmUserModel
	var id = int32(this.FormIntDefault(ctx, "id", 0))
	userModel.ID = id
	gmUser, err := userModel.GetUserById()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, gmUser)
}

// 修改GM用户信息
func (this UserController) UserUpdateAction(ctx *gin.Context) {
	var userModel model.GmUserModel
	var id = int32(this.FormIntDefault(ctx, "id", 0))
	userModel.ID = id
	userModel.Realname = ctx.Request.FormValue("realname")
	userModel.RoleId = int32(this.FormIntDefault(ctx, "role_id", 0))
	userModel.Password = ctx.Request.FormValue("password")
	userModel.Username = ctx.Request.FormValue("username")

	err := userModel.Update()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

//新增GM用户
func (this UserController) UserAddAction(ctx *gin.Context) {
	var userModel model.GmUserModel

	userModel.Realname = ctx.Request.FormValue("realname")
	userModel.RoleId = int32(this.FormIntDefault(ctx, "role_id", 0))
	userModel.Password = ctx.Request.FormValue("password")
	userModel.Username = ctx.Request.FormValue("username")

	id, err := userModel.Insert()
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, id)
}
