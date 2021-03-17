package controller

import (
	"ElasticView/engine/es"

	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
)

type EsMapController struct {
	BaseController
}

func (this EsMapController) ListAction(ctx *gin.Context) {
	esConnect := es.EsConnect{}
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
	res, err := esClinet.GetMappings()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res)
}
