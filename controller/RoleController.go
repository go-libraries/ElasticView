package controller

import (
	"github.com/1340691923/ElasticView/model"
	"github.com/1340691923/ElasticView/platform-basic-libs/response"
	"github.com/1340691923/ElasticView/platform-basic-libs/service/gm_role"

	"github.com/gin-gonic/gin"
)

// GM角色控制器
type RoleController struct {
	BaseController
}

//获取索引的GM 角色
func (this RoleController) RolesAction(ctx *gin.Context) {
	var service gm_role.GmRoleService
	list, err := service.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}

//新增GM角色
func (this RoleController) RolesAddAction(ctx *gin.Context) {

	var model2 model.GmRoleModel

	err := ctx.Bind(&model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var service gm_role.GmRoleService
	id, err := service.Add(model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, map[string]interface{}{"id": id})
}

// 修改GM角色
func (this RoleController) RolesUpdateAction(ctx *gin.Context) {
	var model2 model.GmRoleModel
	err := ctx.Bind(&model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var service gm_role.GmRoleService
	err = service.Update(model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

// 删除GM角色
func (this RoleController) RolesDelAction(ctx *gin.Context) {
	id := this.FormIntDefault(ctx, "id", 0)

	var service gm_role.GmRoleService
	err := service.Delete(id)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

// 获取Gm角色下拉选
func (this RoleController) RoleOptionAction(ctx *gin.Context) {

	var model model.GmRoleModel

	list, err := model.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}
