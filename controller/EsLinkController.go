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

	sql, args, err := db.SqlBuilder.Insert("es_link").SetMap(insertMap).ToSql()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var esLinkList []model.EsLinkModel
	_, err = db.Sqlx.Exec(sql, args...)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, esLinkList)
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

	sql, args, err := db.SqlBuilder.Update("es_link").SetMap(insertMap).Where(db.Eq{"id": esLinkModel.ID}).ToSql()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var esLinkList []model.EsLinkModel
	_, err = db.Sqlx.Exec(sql, args...)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, esLinkList)
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

	sql, args, err := db.SqlBuilder.
		Delete("es_link").
		Where(db.Eq{"id": id}).ToSql()
	log.Println(sql, args)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var esLinkList []model.EsLinkModel
	_, err = db.Sqlx.Exec(sql, args...)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.DeleteSuccess, esLinkList)
}
