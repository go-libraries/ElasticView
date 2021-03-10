package main

import (
	_ "github.com/go-sql-driver/mysql"
	"strconv"

	"ElasticView/application"
	"ElasticView/router"
)


// By 肖文龙
func main(){
	app := application.App{
		ConfigFileDir:  "config",
		ConfigFileName: "config.json",
		AppName:        "Elastic",
	}
	app.InitConfig()
	app.InitLogs()
	app.InitMysql()

	engine := router.Init()
	engine.Run(":"+strconv.Itoa(application.GlobConfig.Port))
}
