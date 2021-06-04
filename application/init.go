package application

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/1340691923/ElasticView/engine/db"
	"github.com/1340691923/ElasticView/engine/logs"
	"github.com/1340691923/ElasticView/model"
	"github.com/1340691923/ElasticView/platform-basic-libs/rbac"
	"github.com/1340691923/ElasticView/platform-basic-libs/util"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Options方法
type NewAppOptions func(app *App)

//App 结构体 启动应用基本配置
type App struct {
	configFileDir,
	configFileName,
	configFileExt,
	appName string
	err error
}

//设置配置文件格式   例如:json,conf 等等
func WithConfigFileExt(configFileExt string) NewAppOptions {
	return func(app *App) {
		app.configFileExt = configFileExt
	}
}

//设置配置文件目录
func WithConfigFileDir(configFileDir string) NewAppOptions {
	return func(app *App) {
		app.configFileDir = configFileDir
	}
}

//设置配置文件名
func WithConfigFileName(configFileName string) NewAppOptions {
	return func(app *App) {
		app.configFileName = configFileName
	}
}

//设置应用名
func WithAppName(appName string) NewAppOptions {
	return func(app *App) {
		app.appName = appName
	}
}

//App 构造方法
func NewApp(opts ...NewAppOptions) *App {
	app := &App{
		configFileDir:  "config",
		configFileName: "config.json",
		appName:        "github.com/1340691923/ElasticView",
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

//初始化配置
func (this *App) InitConfig() *App {

	config := viper.New()
	config.AddConfigPath(this.configFileDir)
	config.SetConfigName(this.configFileName)
	config.SetConfigType(this.configFileExt)
	if err := config.ReadInConfig(); err != nil {
		this.err = err
		return this
	}

	if err := config.Unmarshal(&GlobConfig); err != nil {
		this.err = err
		return this
	}

	config.OnConfigChange(func(e fsnotify.Event) {
		log.Println("检测配置文件更新，应用配置修改中...")
		if err := config.ReadInConfig(); err != nil {
			log.Println("应用配置修改失败", err)
			return
		}
		if err := config.Unmarshal(&GlobConfig); err != nil {
			log.Println("应用配置修改失败", err)
			return
		}
		if err := this.InitLogs().InitMysql().Error(); err != nil {
			log.Println("应用配置修改失败", err)
			return
		}
		log.Println("应用配置修改完成！")
	})

	config.WatchConfig()
	return this
}

//获取配置文件夹
func (this *App) getConfigDir() string {
	return filepath.Join(
		util.GetCurrentDirectory(),
		this.configFileDir,
	)
}

//初始化日志
func (this *App) InitLogs() *App {
	logger := logs.NewLog(
		logs.WithLogPath(GlobConfig.Log.LogDir),
		logs.WithStorageDays(GlobConfig.Log.StorageDays),
	)
	logs.Logger, this.err = logger.InitLog()
	return this
}

//初始化mysql连接
func (this *App) InitMysql() *App {
	config := GlobConfig.Mysql
	db.Sqlx, this.err = db.NewSQLX(
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

//初始化项目启动任务
func (this *App) InitTask() *App {
	esLinkModel := model.EsLinkModel{}
	if err := esLinkModel.FlushEsLinkList(); err != nil {
		this.err = err
		return this
	}
	return this
}

//初始化项目启动任务
func (this *App) InitRbac() *App {
	config := GlobConfig.Mysql

	if err := rbac.Run("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Pwd,
		config.IP,
		config.Port,
		config.DbName),
	); err != nil {
		this.err = err
		return this
	}
	return this
}

//是否有异常
func (this *App) Error() (err error) {
	return this.err
}
