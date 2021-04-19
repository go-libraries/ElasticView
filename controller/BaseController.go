//控制器层
package controller

import (
	"ElasticView/platform-basic-libs/request"
	"ElasticView/platform-basic-libs/response"
)

//父控制器结构体
type BaseController struct {
	response.Response
	request.Request
}
