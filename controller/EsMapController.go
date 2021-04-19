package controller

import (
	"context"

	"ElasticView/engine/es"
	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
)

// Es 映射控制器
type EsMappingController struct {
	BaseController
}

// Es 映射列表
func (this EsMappingController) ListAction(ctx *gin.Context) {
	esConnect := es.EsMapGetProperties{}
	err := ctx.Bind(&esConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esConnect.EsConnectID)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	if esConnect.IndexName == "" {
		res, err := esClinet.GetMappings()
		if err != nil {
			this.Error(ctx, err)
			return
		}
		this.Success(ctx, response.SearchSuccess, res)
	} else {
		res, err := esClinet.(*es.EsClientV6).Client.GetMapping().Index(esConnect.IndexName).Do(ctx)
		if err != nil {
			this.Error(ctx, err)
			return
		}
		this.Success(ctx, response.SearchSuccess, res)
	}

}

/*func (this EsMappingController) GetPropertiesAction(ctx *gin.Context) {
	esConnect := es.EsMapGetProperties{}
	err := ctx.Bind(&esConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esConnect.EsConnectID)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	res, err := esClinet.(*es.EsClientV6).Client.GetMapping().Index("")
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res)
}*/

// 修改映射
func (this EsMappingController) UpdateMappingAction(ctx *gin.Context) {
	updateMapping := es.UpdateMapping{}
	err := ctx.Bind(&updateMapping)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(updateMapping.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	res, err := esClinet.(*es.EsClientV6).Client.PutMapping().
		Index(updateMapping.IndexName).
		Type(updateMapping.TypeName).
		UpdateAllTypes(true).
		BodyJson(updateMapping.Properties).
		Do(context.Background())
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res)
}
