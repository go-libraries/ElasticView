package main

import (
	_ "github.com/go-sql-driver/mysql"

	"ElasticView/application"
)

// By 肖文龙
func main() {
	app := application.NewApp(
		application.WithAppName("ElasticView"),
		application.WithConfigFileDir("config"),
		application.WithConfigFileName("config.json"),
	)

	app.InitConfig().
		InitLogs().
		InitMysql().
		InitTask().
		InitServer()
	select {}
}
