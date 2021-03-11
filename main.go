package main

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"ElasticView/application"
	"ElasticView/router"
)

// By 肖文龙
func main() {
	app := application.App{
		ConfigFileDir:  "config",
		ConfigFileName: "config.json",
		AppName:        "ElasticView",
	}

	app.InitConfig()
	app.InitLogs()
	app.InitMysql()

	port := ":" + strconv.Itoa(application.GlobConfig.Port)
	engine := router.Init()
	engine.Run(port)
}
