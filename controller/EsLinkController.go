package controller

import (
	"encoding/json"
	"log"

	"ElasticView/engine/db"
	"ElasticView/model"
	"ElasticView/platform-basic-libs/request"
	"ElasticView/platform-basic-libs/response"
	"ElasticView/platform-basic-libs/util"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type EsLinkController struct {
	BaseController
}

func (this EsLinkController) ListAction(ctx *gin.Context) {
	sql, args, err := db.SqlBuilder.
		Select("*").
		From("es_link").ToSql()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var esLinkList []model.EsLinkModel
	log.Println(sql, args)
	err = db.Sqlx.Select(&esLinkList, sql, args...)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, esLinkList)
}

func (this EsLinkController) InsertAction(ctx *gin.Context) {

	var esLinkModel model.EsLinkModel
	err = ctx.Bind(&esLinkModel)
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

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this EsLinkController) UpdateAction(ctx *gin.Context) {
	var esLinkModel model.EsLinkModel
	err = ctx.Bind(&esLinkModel)
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

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this EsLinkController) DeleteAction(ctx *gin.Context) {

	err := this.CheckParameter([]request.CheckConfigStruct{
		{request.IdNullError, "id"},
	}, ctx)

	if err != nil {
		this.Error(ctx, err)
		return
	}

	id := util.CtxFormIntDefault(ctx, "id", 0)

	_, err = db.SqlBuilder.
		Delete("es_link").
		Where(db.Eq{"id": id}).RunWith(db.Sqlx).Exec()

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.DeleteSuccess, nil)
}
