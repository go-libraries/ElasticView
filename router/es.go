package router

import (
	. "github.com/1340691923/ElasticView/controller"
	api_config "github.com/1340691923/ElasticView/platform-basic-libs/api_config"
	. "github.com/gofiber/fiber/v2"
)

// ES基础操作 路由
func runEs(app *App) {
	apiRouterConfig := api_config.NewApiRouterConfig()
	const AbsolutePath = "/api/es"
	es := app.Group(AbsolutePath)
	{
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "将索引恢复为可写状态   由于不可抗力，ES禁止写后，默认不会自动恢复",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotRepositoryListAction",
		}, es.(*Group), true, EsController{}.RecoverCanWrite)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "Ping ES",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "PingAction",
		}, es.(*Group), true, EsController{}.PingAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "Es的CAT操作",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "CatAction",
		}, es.(*Group), true, EsController{}.CatAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "运行DSL",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "RunDslAction",
		}, es.(*Group), true, EsController{}.RunDslAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "SQL语法转ES语法",
			Method:       api_config.MethodGet,
			AbsolutePath: AbsolutePath,
			RelativePath: "SqlToDslAction",
		}, es.(*Group), true, EsController{}.SqlToDslAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "一些索引的操作",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "OptimizeAction",
		}, es.(*Group), true, EsController{}.OptimizeAction)

	}
}
