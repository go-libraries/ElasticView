package controller

import (
	"ElasticView/engine/es"
	"ElasticView/platform-basic-libs/request"
	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
)

type EsController struct {
	BaseController
}

//Ping
func (this EsController) PingAction(ctx *gin.Context) {
	esConnect := request.EsConnect{}
	err = ctx.Bind(&esConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClient(esConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	data, _, err := esClinet.Ping()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, data)
}

//Elasticsearch状态
func (this EsController) CatAction(ctx *gin.Context) {
	esCat := request.EsCat{}
	err = ctx.Bind(&esCat)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClient(esCat.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var data interface{}

	switch esCat.Cat {
	case "CatHealth":
		data, err = esClinet.CatHealth()
	case "CatShards":
		data, err = esClinet.CatShards()
	case "CatCount":
		data, err = esClinet.CatCount()
	case "CatAllocation":
		data, err = esClinet.CatAllocation()
	case "CatAliases":
		data, err = esClinet.CatAliases()
	case "CatIndices":
		data, err = esClinet.CatIndices()
	}

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, data)
}
