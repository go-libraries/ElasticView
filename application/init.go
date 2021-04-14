package application

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"ElasticView/engine/db"
	"ElasticView/engine/logs"
	"ElasticView/model"
	"ElasticView/platform-basic-libs/util"

	jsoniter "github.com/json-iterator/go"
)

type App struct {
	ConfigFileDir, ConfigFileName, AppName string
}

func (this *App) InitConfig() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	b, err := ioutil.ReadFile(filepath.Join(this.getConfigDir(), this.ConfigFileName))
	if err != nil {
		panic("config.json 配置文件没找到:" + err.Error())
	}
	err = json.Unmarshal(b, &GlobConfig)
	if err != nil {
		panic("config.json 解析错误:" + err.Error())
	}
	return
}

func (this *App) getConfigDir() string {
	return filepath.Join(
		util.GetCurrentDirectory(),
		this.ConfigFileDir,
	)
}

func (this *App) InitLogs() {
	logs.InitLog(
		GlobConfig.Log.LogDir,
		GlobConfig.Log.StorageDays,
	)
}

func (this *App) InitMysql() {
	config := GlobConfig.Mysql
	db.NewSQLX(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Pwd, config.IP, config.Port, config.DbName),
		config.MaxOpenConns,
		config.MaxIdleConns,
	)
	return
}

func (this *App) InitTask() {
	esLinkModel := model.EsLinkModel{}
	if err := esLinkModel.FlushEsLinkList(); err != nil {
		panic(err)
	}
	return
}
