//中间件层
package middleware

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/1340691923/ElasticView/engine/logs"
	"github.com/1340691923/ElasticView/platform-basic-libs/jwt"
	"github.com/1340691923/ElasticView/platform-basic-libs/my_error"
	"github.com/1340691923/ElasticView/platform-basic-libs/response"
	"github.com/1340691923/ElasticView/platform-basic-libs/service/gm_user"
	"github.com/1340691923/ElasticView/platform-basic-libs/util"

	"github.com/gin-gonic/gin"
)

var res response.Response

// 处理跨域请求,支持options访问
func Cors(ctx *gin.Context) {

	method := ctx.Request.Method
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, X-Token")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	// 放行所有OPTIONS方法，因为有的模板是要请求两次的
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	ctx.Next()
}

func JwtMiddleware(ctx *gin.Context) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			//打印调用栈信息
			buf := make([]byte, 2048)
			logger := logs.Logger.Sugar()
			n := runtime.Stack(buf, false)
			stackInfo := fmt.Sprintf("%s", buf[:n])
			logger.Errorf("panic stack info %s", stackInfo)
			logger.Errorf("--->JwtMiddleware Error:", r)
		}
	}()

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type") //header的类型
	ctx.Header("Access-Control-Allow-Headers", "x-token")      //header的类型
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if strings.ToUpper(ctx.Request.Method) == "OPTIONS" {
		ctx.Abort()
		return
	}

	var claims *jwt.Claims
	if util.GetToken(ctx) == "" {
		err = my_error.NewBusiness(TOKEN_ERROR, INVALID_PARAMS)
		res.Error(ctx, err)
		ctx.Abort()
		return
	} else {
		token := util.GetToken(ctx)

		var service gm_user.GmUserService
		claims, err = jwt.ParseToken(token)
		if err != nil {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_FAIL)
			res.Error(ctx, err)
			ctx.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			res.Error(ctx, err)
			ctx.Abort()
			return
		} else if !service.IsExitUser(claims) {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			res.Error(ctx, err)
			ctx.Abort()
			return
		}
	}

	ctx.Next()
}
