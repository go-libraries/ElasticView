package controller

import (
	"github.com/1340691923/ElasticView/engine/es"
	"github.com/1340691923/ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

// Es 任务控制器
type TaskController struct {
	BaseController
}

// 任务列表
func (this TaskController) ListAction(ctx *gin.Context) {
	taskListReq := es.TaskList{}
	err := ctx.Bind(&taskListReq)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(taskListReq.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	tasksListService := esClinet.(*es.EsClientV6).Client.TasksList().Detailed(true)

	tasksListResponse, err := tasksListService.Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	taskListRes := map[string]*elastic.TaskInfo{}

	for _, node := range tasksListResponse.Nodes {
		for taskId, taskInfo := range node.Tasks {
			taskListRes[taskId] = taskInfo
		}
	}

	this.Success(ctx, response.SearchSuccess, taskListRes)
	return
}

// 取消任务
func (this TaskController) CancelAction(ctx *gin.Context) {
	cancelTask := es.CancelTask{}
	err := ctx.Bind(&cancelTask)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(cancelTask.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	res, err := esClinet.(*es.EsClientV6).Client.TasksCancel().TaskId(cancelTask.TaskID).Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res)
}
