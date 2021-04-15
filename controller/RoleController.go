package controller

import (
	"ElasticView/model"
	"ElasticView/platform-basic-libs/response"
	"ElasticView/platform-basic-libs/service/gm_role"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (this RoleController) RolesAction(ctx *gin.Context) {
	var service gm_role.GmRoleService
	list, err := service.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}

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

func (this RoleController) RoleOptionAction(ctx *gin.Context) {

	var model model.GmRoleModel

	list, err := model.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}
