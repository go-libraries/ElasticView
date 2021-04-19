package main

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"ElasticView/application"
	"ElasticView/router"
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
		panic(err)
	}

	port := ":" + strconv.Itoa(application.GlobConfig.Port)
	engine := router.Init()

	err = engine.Run(port)
	if err != nil {
		panic(err)
	}

}
