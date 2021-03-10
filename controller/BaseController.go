package controller

import (
	"ElasticView/platform-basic-libs/request"
	"ElasticView/platform-basic-libs/response"
)

type BaseController struct {
	response.Response
	request.Request
}

var err error
