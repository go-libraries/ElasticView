package main

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/1340691923/ElasticView/application"
	"github.com/1340691923/ElasticView/engine/logs"
	"github.com/1340691923/ElasticView/router"
	"go.uber.org/zap"
)

// By 肖文龙
func main() {
	app := application.NewApp(
		application.WithAppName("ElasticView"),
		application.WithConfigFileDir("config"),
		application.WithConfigFileName("config"),
		application.WithConfigFileExt("json"),
	)

	err := app.InitConfig().
		InitLogs().
		InitMysql().
		InitTask().
		Error()
	if err != nil {
		logs.Logger.Error("ElasticView 初始化失败", zap.String("err.Error()", err.Error()))
		panic(err)
	}

	port := ":" + strconv.Itoa(application.GlobConfig.Port)
	engine := router.Init()

	err = engine.Run(port)
	if err != nil {
		logs.Logger.Error("ElasticView http服务器启动失败:", zap.String("err.Error()", err.Error()))
		panic(err)
	}

}
