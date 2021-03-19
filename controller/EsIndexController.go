package controller

import (
	"context"

	"ElasticView/engine/es"
	"ElasticView/engine/logs"
	"ElasticView/platform-basic-libs/my_error"
	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
)

type EsIndexController struct {
	BaseController
}

//创建索引
func (this EsIndexController) CreateAction(ctx *gin.Context) {
	esIndexInfo := es.EsIndexInfo{}
	err = ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClient(esIndexInfo.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	if esIndexInfo.IndexName == "" {
		this.Error(ctx, my_error.NewBusiness(es.ParmasNullError, es.IndexNameNullError))
		return
	}
	logs.Logger.Sugar().Infof("esIndexInfo.Types", esIndexInfo.Types)
	if esIndexInfo.Types == "update" {
		res, err := esClinet.IndexPutSettings(esIndexInfo.IndexName, esIndexInfo.Settings)
		if err != nil {
			this.Error(ctx, err)
			return
		}
		this.Success(ctx, response.OperateSuccess, res)
		return
	} else {
		res, err := esClinet.(*es.EsClientV6).Client.CreateIndex(esIndexInfo.IndexName).BodyJson(map[string]interface{}{
			"settings": esIndexInfo.Settings,
		}).Do(context.Background())
		if err != nil {
			this.Error(ctx, err)
			return
		}
		this.Success(ctx, response.OperateSuccess, res)
		return
	}
	return
}

//创建索引
func (this EsIndexController) GetSettingsAction(ctx *gin.Context) {
	esIndexInfo := es.EsIndexInfo{}
	err = ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClient(esIndexInfo.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	if esIndexInfo.IndexName == "" {
		this.Error(ctx, my_error.NewBusiness(es.ParmasNullError, es.IndexNameNullError))
		return
	}

	res, err := esClinet.(*es.EsClientV6).Client.IndexGetSettings(esIndexInfo.IndexName).Do(context.Background())
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, res[esIndexInfo.IndexName].Settings)
	return
}

//获取别名
func (this EsIndexController) GetAliasAction(ctx *gin.Context) {
	esIndexInfo := es.EsIndexInfo{}
	err = ctx.Bind(&esIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClient(esIndexInfo.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	if esIndexInfo.IndexName == "" {
		this.Error(ctx, my_error.NewBusiness(es.ParmasNullError, es.IndexNameNullError))
		return
	}

	aliasRes, err := esClinet.(*es.EsClientV6).Client.Aliases().Index(esIndexInfo.IndexName).Do(context.TODO())

	this.Success(ctx, response.OperateSuccess, aliasRes.Indices[esIndexInfo.IndexName].Aliases)
	return
}

func (this EsIndexController) ReindexAction(ctx *gin.Context) {
	esReIndexInfo := es.EsReIndexInfo{}
	err = ctx.Bind(&esReIndexInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClient(esReIndexInfo.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	reindex := esClinet.(*es.EsClientV6).Client.Reindex()
	urlValues := esReIndexInfo.UrlValues
	if urlValues.WaitForActiveShards != "" {
		reindex = reindex.WaitForActiveShards(urlValues.WaitForActiveShards)
	}
	if urlValues.Slices != 0 {
		reindex = reindex.Slices(urlValues.Slices)
	}
	if urlValues.Refresh != "" {
		reindex = reindex.Refresh(urlValues.Refresh)
	}
	if urlValues.Timeout != "" {
		reindex = reindex.Timeout(urlValues.Refresh)
	}
	if urlValues.RequestsPerSecond != 0 {
		reindex = reindex.RequestsPerSecond(urlValues.RequestsPerSecond)
	}
	if urlValues.WaitForCompletion != nil {
		reindex = reindex.WaitForCompletion(*urlValues.WaitForCompletion)
	}

	res, err := reindex.Body(esReIndexInfo.Body).Do(context.Background())
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res)
}

func (this EsIndexController) ReindexListAction(ctx *gin.Context) {

}

func (this EsIndexController) IndexNamesAction(ctx *gin.Context) {
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
	indexNames, err := esClinet.(*es.EsClientV6).Client.IndexNames()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, indexNames)
	return
}
