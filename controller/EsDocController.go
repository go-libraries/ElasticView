package controller

import (
	"fmt"

	"ElasticView/engine/es"
	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

type EsDocController struct {
	BaseController
}

//删除文档数据
func (this EsDocController) DeleteRowByIDAction(ctx *gin.Context) {
	esDocDeleteRowByID := es.EsDocDeleteRowByID{}
	err := ctx.Bind(&esDocDeleteRowByID)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esDocDeleteRowByID.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	res, err := esClinet.(*es.EsClientV6).Client.PerformRequest(ctx, elastic.PerformRequestOptions{
		Method: "DELETE",
		Path:   fmt.Sprintf("/%s/%s/%s", esDocDeleteRowByID.IndexName, esDocDeleteRowByID.Type, esDocDeleteRowByID.ID),
		Body:   nil,
	})

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res)
}

//删除文档数据
func (this EsDocController) UpdateByIDAction(ctx *gin.Context) {
	esDocUpdateByID := es.EsDocUpdateByID{}
	err := ctx.Bind(&esDocUpdateByID)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esDocUpdateByID.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	res, err := esClinet.(*es.EsClientV6).Client.Update().Index(esDocUpdateByID.Index).Type(esDocUpdateByID.Type).Id(esDocUpdateByID.ID).
		Doc(esDocUpdateByID.JSON).Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res)
}

func (this EsDocController) InsertAction(ctx *gin.Context) {
	esDocUpdateByID := es.EsDocUpdateByID{}
	err := ctx.Bind(&esDocUpdateByID)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esDocUpdateByID.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	res, err := esClinet.(*es.EsClientV6).Client.Index().
		Index(esDocUpdateByID.Index).
		Type(esDocUpdateByID.Type).BodyJson(esDocUpdateByID.JSON).Do(ctx)
	if err != nil {
		return
	}
	this.Success(ctx, response.OperateSuccess, res)
}
