package application

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"ElasticView/engine/db"
	"ElasticView/engine/logs"
	"ElasticView/model"
	"ElasticView/platform-basic-libs/util"
	"ElasticView/router"

	"github.com/fvbock/endless"
	jsoniter "github.com/json-iterator/go"
)

type NewAppOptions func(app *App)

//App 结构体
type App struct {
	configFileDir,
	configFileName,
	appName string
}

func WithConfigFileDir(configFileDir string) NewAppOptions {
	return func(app *App) {
		app.configFileDir = configFileDir
	}
}

func WithConfigFileName(configFileName string) NewAppOptions {
	return func(app *App) {
		app.configFileName = configFileName
	}
}
func WithAppName(appName string) NewAppOptions {
	return func(app *App) {
		app.appName = appName
	}
}

func NewApp(opts ...NewAppOptions) *App {
	app := &App{
		configFileDir:  "config",
		configFileName: "config.json",
		appName:        "ElasticView",
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

func (this *App) InitConfig() *App {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	b, err := ioutil.ReadFile(filepath.Join(this.getConfigDir(), this.configFileName))
	if err != nil {
		panic("config.json 配置文件没找到:" + err.Error())
	}
	err = json.Unmarshal(b, &GlobConfig)
	if err != nil {
		panic("config.json 解析错误:" + err.Error())
	}
	return this
}

func (this *App) getConfigDir() string {
	return filepath.Join(
		util.GetCurrentDirectory(),
		this.configFileDir,
	)
}

func (this *App) InitLogs() *App {
	logs.InitLog(
		GlobConfig.Log.LogDir,
		GlobConfig.Log.StorageDays,
	)
	return this
}

func (this *App) InitMysql() *App {
	config := GlobConfig.Mysql
	db.NewSQLX(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			config.Username,
			config.Pwd,
			config.IP,
			config.Port,
			config.DbName),
		config.MaxOpenConns,
		config.MaxIdleConns,
	)
	return this
}

func (this *App) InitTask() *App {
	esLinkModel := model.EsLinkModel{}
	if err := esLinkModel.FlushEsLinkList(); err != nil {
		panic(err)
	}
	return this
}

func (this *App) InitServer() *App {
	port := ":" + strconv.Itoa(GlobConfig.Port)
	engine := router.Init()
	server := endless.NewServer(port, engine)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	return this
}
