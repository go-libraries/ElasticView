package controller

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"ElasticView/engine/es"
	"ElasticView/platform-basic-libs/request"
	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
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

func (this EsController) RunDslAction(ctx *gin.Context) {
	esRest := request.EsRest{}
	err = ctx.Bind(&esRest)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClient(esRest.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	res, err := esClinet.(*es.EsClientV6).Client.PerformRequest(context.TODO(), elastic.PerformRequestOptions{
		Method: esRest.Method,
		Path:   esRest.Path,
		Body:   esRest.Body,
	})

	if err != nil {
		this.Error(ctx, err)
		return
	}

	if res.StatusCode != 200 {
		this.Output(ctx, map[string]interface{}{
			"code": res.StatusCode,
			"msg":  errors.New(fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode))),
			"data": res.Body,
		})
		return
	}

	this.Success(ctx, response.SearchSuccess, res.Body)
}
