package controller

import (
	"encoding/json"

	"github.com/1340691923/ElasticView/engine/db"
	"github.com/1340691923/ElasticView/engine/es"
	"github.com/1340691923/ElasticView/engine/logs"
	"github.com/1340691923/ElasticView/model"
	"github.com/1340691923/ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// Es 连接管理控制器
type EsLinkController struct {
	BaseController
}

// 获取Es连接列表
func (this EsLinkController) ListAction(ctx *gin.Context) {
	logs.Logger.Sugar().Infof("logs.Logger.Sugar: %v\n", "111")
	getByLocal := ctx.Request.FormValue("getByLocal")

	if getByLocal == "1" {
		this.Success(ctx, response.SearchSuccess, model.EsLinkList)
		return
	}

	esLinkModel := model.EsLinkModel{}

	list, err := esLinkModel.GetListAction()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}

// 新增Es连接
func (this EsLinkController) InsertAction(ctx *gin.Context) {

	var esLinkModel model.EsLinkModel
	err := ctx.Bind(&esLinkModel)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	esB, err := json.Marshal(esLinkModel)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	insertMap := map[string]interface{}{}
	err = json.Unmarshal(esB, &insertMap)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	delete(insertMap, "created")
	delete(insertMap, "updated")

	_, err = db.SqlBuilder.
		Insert("es_link").
		SetMap(insertMap).
		RunWith(db.Sqlx).
		Exec()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	err = esLinkModel.FlushEsLinkList()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

// 修改Es连接信息
func (this EsLinkController) UpdateAction(ctx *gin.Context) {
	var esLinkModel model.EsLinkModel
	err := ctx.Bind(&esLinkModel)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esB, err := json.Marshal(esLinkModel)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	insertMap := map[string]interface{}{}
	err = json.Unmarshal(esB, &insertMap)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	delete(insertMap, "id")
	delete(insertMap, "created")
	delete(insertMap, "updated")

	_, err = db.SqlBuilder.
		Update("es_link").
		SetMap(insertMap).
		Where(db.Eq{"id": esLinkModel.ID}).
		RunWith(db.Sqlx).
		Exec()
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esCache := es.NewEsCache()
	esCache.Rem(int(esLinkModel.ID))

	err = esLinkModel.FlushEsLinkList()
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
}

// 删除es连接
func (this EsLinkController) DeleteAction(ctx *gin.Context) {

	var req struct {
		Id int `json:"id"`
	}

	err := ctx.Bind(&req)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	_, err = db.SqlBuilder.
		Delete("es_link").
		Where(db.Eq{"id": req.Id}).RunWith(db.Sqlx).Exec()

	if err != nil {
		this.Error(ctx, err)
		return
	}

	esCache := es.NewEsCache()
	esCache.Rem(req.Id)
	esLinkModel := model.EsLinkModel{}
	err = esLinkModel.FlushEsLinkList()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.DeleteSuccess, nil)
}
