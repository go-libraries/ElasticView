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
