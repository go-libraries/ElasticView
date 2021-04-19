package controller

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"ElasticView/engine/es"
	"ElasticView/model"
	"ElasticView/platform-basic-libs/jwt"
	"ElasticView/platform-basic-libs/response"
	"ElasticView/platform-basic-libs/service/es_optimize"

	"github.com/cch123/elasticsql"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

//Es 基本操作
type EsController struct {
	BaseController
}

// Ping
func (this EsController) PingAction(ctx *gin.Context) {
	esConnect := es.EsConnect{}
	err := ctx.Bind(&esConnect)
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

// Es 的CAT API
func (this EsController) CatAction(ctx *gin.Context) {

	esCat := es.EsCat{}
	err := ctx.Bind(&esCat)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esCat.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var data interface{}

	switch esCat.Cat {
	case "CatHealth":
		data, err = esClinet.(*es.EsClientV6).Client.CatHealth().Human(true).Do(ctx)
	case "CatShards":
		data, err = esClinet.(*es.EsClientV6).Client.CatShards().Human(true).Do(ctx)
	case "CatCount":
		data, err = esClinet.(*es.EsClientV6).Client.CatCount().Human(true).Do(ctx)
	case "CatAllocation":
		data, err = esClinet.(*es.EsClientV6).Client.CatAllocation().Human(true).Do(ctx)
	case "CatAliases":
		data, err = esClinet.(*es.EsClientV6).Client.CatAliases().Human(true).Do(ctx)
	case "CatIndices":
		if esCat.IndexBytesFormat != "" {
			data, err = esClinet.(*es.EsClientV6).Client.CatIndices().Human(true).Bytes(esCat.IndexBytesFormat).Do(ctx)
		} else {
			data, err = esClinet.(*es.EsClientV6).Client.CatIndices().Human(true).Do(ctx)
		}
	case "CatSegments":
		data, err = esClinet.(*es.EsClientV6).Client.IndexSegments().Human(true).Do(ctx)
	case "CatStats":
		data, err = esClinet.(*es.EsClientV6).Client.ClusterStats().Human(true).Do(ctx)
	}

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, data)
}

func (this EsController) RunDslAction(ctx *gin.Context) {

	esRest := es.EsRest{}
	err := ctx.Bind(&esRest)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esRest.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esRest.Method = strings.ToUpper(esRest.Method)
	if esRest.Method == "GET" {
		c, err := jwt.ParseToken(ctx.GetHeader("X-Token"))
		if err != nil {
			this.Error(ctx, err)
			return
		}

		gmDslHistoryModel := model.GmDslHistoryModel{
			Uid:    int(c.ID),
			Method: esRest.Method,
			Path:   esRest.Path,
			Body:   esRest.Body,
		}

		err = gmDslHistoryModel.Insert()

		if err != nil {
			this.Error(ctx, err)
			return
		}
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

	if res.StatusCode != 200 && res.StatusCode != 201 {
		this.Output(ctx, map[string]interface{}{
			"code": res.StatusCode,
			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode)),
			"data": res.Body,
		})
		return
	}

	this.Success(ctx, response.OperateSuccess, res.Body)
}

// SQL 转换为 DSL
func (this EsController) SqlToDslAction(ctx *gin.Context) {
	sql := ctx.Request.FormValue("sql")
	dsl, table, err := elasticsql.ConvertPretty(sql)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "转换成功!", map[string]interface{}{
		"dsl":       dsl,
		"tableName": table,
	})
}

// 一些索引的操作
func (this EsController) OptimizeAction(ctx *gin.Context) {
	esOptimize := es.EsOptimize{}
	err := ctx.Bind(&esOptimize)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esOptimize.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	optimize := es_optimize.OptimizeFactory(esOptimize.Command)

	if optimize == nil {
		this.Error(ctx, errors.New("不支持该指令"))
		return
	}
	if esOptimize.IndexName != "" {
		optimize.SetIndexName(esOptimize.IndexName)
	}
	err = optimize.Do(esClinet.(*es.EsClientV6).Client)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

// 将索引恢复为可写状态   由于不可抗力，ES禁止写后，默认不会自动恢复
func (this EsController) RecoverCanWrite(ctx *gin.Context) {
	esConnect := es.EsConnectID{}
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

	res, err := esClinet.(*es.EsClientV6).Client.PerformRequest(ctx, elastic.PerformRequestOptions{
		Method: "PUT",
		Path:   "/_settings",
		Body: map[string]interface{}{
			"index": map[string]interface{}{
				"blocks": map[string]interface{}{
					"read_only_allow_delete": "false",
				},
			},
		},
	})

	if res.StatusCode != 200 && res.StatusCode != 201 {
		this.Output(ctx, map[string]interface{}{
			"code": res.StatusCode,
			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode)),
			"data": res.Body,
		})
		return
	}

	this.Success(ctx, response.OperateSuccess, res.Body)
}
