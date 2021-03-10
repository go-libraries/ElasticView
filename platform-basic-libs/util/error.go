package util

import (
	"database/sql"
	"github.com/garyburd/redigo/redis"
	"io"
)

const (
	ErrNotError = 0

	ErrNoData       = 96
	ErrParamInvalid = 97
	ErrSignFailed   = 98
	ErrUnknown      = 99

	ErrNoUpdate = 201

	ErrShouldLogin      = 210
	ErrPermissionDenied = 211
	ErrPwdNotMatch      = 212
	ErrLoginFailed      = 213

	ErrPayDuplicate = -224
	ErrPayClosed    = -228

	ErrSerial       = 301
	ErrSiteNotFound = 302
)

var errMsg = map[int]string{
	ErrNotError: "成功",

	ErrNoData:       "没有数据",
	ErrParamInvalid: "参数错误",
	ErrSignFailed:   "验证签名错误",
	ErrUnknown:      "未知错误",

	ErrNoUpdate: "无更新",

	ErrShouldLogin:      "登录过期",
	ErrPermissionDenied: "权限不足",
	ErrPwdNotMatch:      "账号密码不匹配",
	ErrLoginFailed:      "验证登录失败",

	ErrPayDuplicate: "订单已处理",
	ErrPayClosed:    "支付未开启",

	ErrSerial:       "序号落后",
	ErrSiteNotFound: "互推位不存在",
}

// ErrCode 状态信息
type ErrCode = int
type Extra = map[string]interface{}

// WriteErr 输出统一的对外格式
func WriteErr(w io.Writer, code ErrCode, extra ...Extra) (int, error) {
	var d map[string]interface{}
	if extra == nil || extra[0] == nil {
		d = map[string]interface{}{}
	} else {
		d = extra[0]
	}
	d["code"] = code
	var ok bool
	d["msg"], ok = errMsg[code]
	if !ok {
		d["msg"] = "unknown code"
	}
	return WriteJSON(w, d)
}

type Error int

func (e Error) Error() string {
	return errMsg[int(e)]
}

func FilterMysqlNilErr(err error) bool {
	if err != nil && err != sql.ErrNoRows {
		return true
	}
	return false
}

func FilterRedisNilErr(err error) bool {
	if err != nil && err != redis.ErrNil {
		return true
	}
	return false
}
