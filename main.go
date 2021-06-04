package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/1340691923/ElasticView/platform-basic-libs/util"
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
		InitRbac().
		Error()
	if err != nil {
		logs.Logger.Error("ElasticView 初始化失败", zap.String("err.Error()", err.Error()))
		panic(err)
	}

	port := ":" + strconv.Itoa(application.GlobConfig.Port)
	appServer := router.Init()

	go func() {
		if err := appServer.Listen(port); err != nil {
			logs.Logger.Error("ElasticView http服务启动失败:", zap.String("err.Error()", err.Error()))
			log.Panic(err)
		}
	}()
	util.OpenWinBrowser(fmt.Sprintf("%s%s", "http://127.0.0.1", port))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	logs.Logger.Info("ElasticView http服务停止中...")
	//这里进行任务释放操作
	/*err = appServer.Shutdown()
	if err != nil {
		log.Println("err", err)
		logs.Logger.Error("ElasticView http服务停止失败:", zap.String("err.Error()", err.Error()))
	}*/
	logs.Logger.Info("ElasticView http服务停止成功...")
}
