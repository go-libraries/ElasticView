package response

import (
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"ElasticView/engine/logs"
	."ElasticView/platform-basic-libs/my_error"
	"ElasticView/platform-basic-libs/util"

	jsoniter "github.com/json-iterator/go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = 500
)

const (
	SearchSuccess  = "查询成功"
	DeleteSuccess  = "删除成功"
	OperateSuccess = "操作成功"
	LogoutSuccess  = "注销成功"
)

func (this *Response) JsonDealErr(err error) string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := json.Marshal(this.DealErr(err))
	return util.BytesToStr(b)
}

//trace
func (this *Response) DealErr(err error) (errorTrace []string) {
	errorTrace = append(errorTrace, err.Error())
	if err != nil {
		for i := 1; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			f := runtime.FuncForPC(pc)
			if f.Name() != "runtime.main" && f.Name() != "runtime.goexit" && !strings.Contains(file, "Response.go") {
				errStrings := "文件名:" + file + ",行数:" + strconv.Itoa(line) + ",函数名:" + f.Name()
				errorTrace = append(errorTrace, errStrings)
			}
		}
	}
	return errorTrace
}

//正确信息
func (this *Response) Success(ctx *gin.Context, msg string, data interface{}) {
	this.send(ctx.Writer, msg, SUCCESS, data)
}

//错误信息
func (this *Response) Error(ctx *gin.Context, err error) {
	errorTrace := this.getTrace(err)

	myErr := ErrorToErrorCode(err)

	logs.Logger.Error("Error",  zap.Strings("err", this.DealErr(myErr)))

	this.send(ctx.Writer, myErr.Error(), myErr.Code(), errorTrace)
}

//输出
func (this *Response) send(ctx io.Writer, msg string, code int, data interface{}) {
	var res Response
	res.Code = code
	res.Msg = msg
	res.Data = data
	util.WriteJSON(ctx, res)
	return
}

//输出
func (this *Response) Output(ctx gin.Context, data interface{}) {
	util.WriteJSON(ctx.Writer, data)
	return
}

//得到trace信息
func (this *Response) getTrace(err error) []string {
	goEnv := os.Getenv("GO_ENV")
	errorTrace := []string{}
	if goEnv == "product" {
		errorTrace = this.DealErr(err)
	}
	return errorTrace
}

//处理异常（业务异常和默认异常）
func ErrorToErrorCode(err error) *MyError {
	if err == nil {
		return nil
	}
	errorCode, ok := err.(*MyError)

	if ok {
		return errorCode
	}
	return NewError(err.Error(), ERROR).(*MyError)
}

func (this *Response) ReturnValOrNull(value, empty interface{}) interface{} {
	var vValue = reflect.ValueOf(value)
	if value == nil || (vValue.Kind() == reflect.Slice && vValue.Len() == 0) {
		return empty
	}
	return value
}

func (this *Response) SliceReturnValOrNull(value []string, empty interface{}) interface{} {
	if value == nil || len(value) == 0 {
		return empty
	}
	return value
}
