package controller

import (
	"ElasticView/engine/es"
	"ElasticView/platform-basic-libs/my_error"

	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
)

type EsMappingController struct {
	BaseController
}

func (this EsMappingController) ListAction(ctx *gin.Context) {
	esConnect := es.EsConnectID{}
	err = ctx.Bind(&esConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientByID(esConnect.EsConnectID)
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

func (this EsMappingController) PatchGetMappingAction(ctx *gin.Context) {

}

func (this EsMappingController) UpdateMappingAction(ctx *gin.Context) {
	esMappingInfo := es.EsMappingInfo{}
	err = ctx.Bind(&esMappingInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	_, err := es.GetEsClientByID(esMappingInfo.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	if esMappingInfo.IndexName == "" {
		this.Error(ctx, my_error.NewBusiness(es.ParmasNullError, es.IndexNameNullError))
		return
	}

}
