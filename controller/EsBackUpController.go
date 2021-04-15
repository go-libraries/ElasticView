package controller

import (
	"errors"
	"fmt"
	"strings"

	"ElasticView/engine/es"
	"ElasticView/engine/logs"
	"ElasticView/platform-basic-libs/my_error"
	"ElasticView/platform-basic-libs/response"
	"ElasticView/platform-basic-libs/service/es_settings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic"
)

type EsBackUpController struct {
	BaseController
}

func (this EsBackUpController) SnapshotRepositoryListAction(ctx *gin.Context) {
	esSnapshotInfo := es.EsSnapshotInfo{}
	err := ctx.Bind(&esSnapshotInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esSnapshotInfo.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	clusterSettings, err := es_settings.NewSettings(esClinet.(*es.EsClientV6).Client)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	pathRepo := clusterSettings.GetPathRepo()

	if len(pathRepo) == 0 {
		this.Error(ctx, my_error.NewError(`path.repo没有设置`, 199999))
		return
	}

	res, err := esClinet.(*es.EsClientV6).Client.SnapshotGetRepository(esSnapshotInfo.SnapshotInfoList...).Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	type tmp struct {
		Name                   string `json:"name"`
		Type                   string `json:"type"`
		Location               string `json:"location"`
		Compress               string `json:"compress"`
		MaxRestoreBytesPerSec  string `json:"max_restore_bytes_per_sec"`
		MaxSnapshotBytesPerSec string `json:"max_snapshot_bytes_per_sec"`
		ChunkSize              string `json:"chunk_size"`
		Readonly               string `json:"readonly"`
	}
	list := []tmp{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for name, settings := range res {
		var t tmp
		t.Type = settings.Type
		t.Name = name
		b, err := json.Marshal(settings.Settings)
		if err != nil {
			logs.Logger.Sugar().Errorf("err", err)
			continue
		}
		err = json.Unmarshal(b, &t)
		if err != nil {
			logs.Logger.Sugar().Errorf("err", err)
			continue
		}
		list = append(list, t)
	}

	this.Success(ctx, response.SearchSuccess, map[string]interface{}{
		"list":     list,
		"res":      res,
		"pathRepo": pathRepo,
	})
}

func (this EsBackUpController) SnapshotCreateRepositoryAction(ctx *gin.Context) {
	snapshotCreateRepository := es.SnapshotCreateRepository{}
	err := ctx.Bind(&snapshotCreateRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(snapshotCreateRepository.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	clusterSettings, err := es_settings.NewSettings(esClinet.(*es.EsClientV6).Client)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	pathRepo := clusterSettings.GetPathRepo()
	getAllowedUrls := clusterSettings.GetAllowedUrls()

	settings := map[string]interface{}{}

	if snapshotCreateRepository.Compress != "" {
		compress := snapshotCreateRepository.Compress
		settings["compress"] = compress
	}

	if snapshotCreateRepository.MaxRestoreBytesPerSec != "" {
		settings["max_restore_bytes_per_sec"] = snapshotCreateRepository.MaxRestoreBytesPerSec
	}

	if snapshotCreateRepository.MaxSnapshotBytesPerSec != "" {
		settings["max_snapshot_bytes_per_sec"] = snapshotCreateRepository.MaxSnapshotBytesPerSec
	}

	if snapshotCreateRepository.Readonly != "" {
		settings["readonly"] = snapshotCreateRepository.Readonly
	}

	if snapshotCreateRepository.ChunkSize != "" {
		settings["chunk_size"] = snapshotCreateRepository.ChunkSize
	}

	switch snapshotCreateRepository.Type {
	case "fs":
		if len(pathRepo) == 0 {
			this.Error(ctx, errors.New("请先设置 path.repo"))
			return
		}
		settings["location"] = snapshotCreateRepository.Location
	case "url":
		if len(getAllowedUrls) == 0 {
			this.Error(ctx, errors.New("请先设置 allowed_urls"))
			return
		}
		settings["url"] = snapshotCreateRepository.Location
	default:
		this.Error(ctx, errors.New("无效的type"))
		return
	}

	_, err = esClinet.(*es.EsClientV6).Client.SnapshotCreateRepository(snapshotCreateRepository.Repository).Type(snapshotCreateRepository.Type).Settings(
		settings,
	).Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this EsBackUpController) CleanupeRepositoryAction(ctx *gin.Context) {
	cleanupeRepository := es.CleanupeRepository{}
	err := ctx.Bind(&cleanupeRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(cleanupeRepository.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	res, err := esClinet.(*es.EsClientV6).Client.PerformRequest(ctx, elastic.PerformRequestOptions{
		Method: "POST",
		Path:   fmt.Sprintf("/_snapshot/%s/_cleanup", cleanupeRepository.Repository),
	})
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, res.Body)
}

func (this EsBackUpController) SnapshotDeleteRepositoryAction(ctx *gin.Context) {
	snapshotDeleteRepository := es.SnapshotDeleteRepository{}
	err := ctx.Bind(&snapshotDeleteRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(snapshotDeleteRepository.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	_, err = esClinet.(*es.EsClientV6).Client.SnapshotDeleteRepository(snapshotDeleteRepository.Repository).Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this EsBackUpController) CreateSnapshotAction(ctx *gin.Context) {
	createSnapshot := es.CreateSnapshot{}
	err := ctx.Bind(&createSnapshot)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(createSnapshot.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	snapshotCreateService := esClinet.(*es.EsClientV6).Client.
		SnapshotCreate(createSnapshot.RepositoryName, createSnapshot.SnapshotName)

	if createSnapshot.Wait != nil {
		snapshotCreateService.WaitForCompletion(*createSnapshot.Wait)
	}

	settings := es.Json{}

	if len(createSnapshot.IndexList) > 0 {
		settings["indices"] = strings.Join(createSnapshot.IndexList, ",")
	}

	if createSnapshot.IgnoreUnavailable != nil {
		settings["indices"] = *createSnapshot.IgnoreUnavailable
	}

	if createSnapshot.Partial != nil {
		settings["partial"] = *createSnapshot.Partial
	}
	if createSnapshot.IncludeGlobalState != nil {
		settings["include_global_state"] = *createSnapshot.IncludeGlobalState
	}

	res, err := snapshotCreateService.BodyJson(settings).Do(ctx)

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, res)
}

func (this EsBackUpController) SnapshotListAction(ctx *gin.Context) {
	snapshotList := es.SnapshotList{}
	err := ctx.Bind(&snapshotList)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(snapshotList.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	if snapshotList.Repository == "" {
		this.Error(ctx, errors.New("请先选择快照存储库"))
		return
	}

	res, err := esClinet.(*es.EsClientV6).Client.PerformRequest(ctx, elastic.PerformRequestOptions{
		Method: "GET",
		Path:   fmt.Sprintf("/_cat/snapshots/%s", snapshotList.Repository),
	})

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, res.Body)
}

func (this EsBackUpController) SnapshotDeleteAction(ctx *gin.Context) {
	snapshotDelete := es.SnapshotDelete{}
	err := ctx.Bind(&snapshotDelete)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(snapshotDelete.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	_, err = esClinet.(*es.EsClientV6).Client.
		SnapshotDelete(snapshotDelete.Repository, snapshotDelete.Snapshot).Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this EsBackUpController) SnapshotDetailAction(ctx *gin.Context) {
	snapshotDetail := es.SnapshotDetail{}
	err := ctx.Bind(&snapshotDetail)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	esClinet, err := es.GetEsClientV6ByID(snapshotDetail.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	res, err := esClinet.(*es.EsClientV6).Client.PerformRequest(ctx, elastic.PerformRequestOptions{
		Method: "GET",
		Path:   fmt.Sprintf("/_snapshot/%s/%s", snapshotDetail.Repository, snapshotDetail.Snapshot),
	})
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res.Body)
}

func (this EsBackUpController) SnapshotRestoreAction(ctx *gin.Context) {
	snapshotRestore := es.SnapshotRestore{}
	err := ctx.Bind(&snapshotRestore)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(snapshotRestore.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	snapshotRestoreService := esClinet.(*es.EsClientV6).Client.SnapshotRestore(snapshotRestore.RepositoryName, snapshotRestore.SnapshotName)

	if snapshotRestore.Wait != nil {
		snapshotRestoreService.WaitForCompletion(*snapshotRestore.Wait)
	}

	if snapshotRestore.IgnoreUnavailable != nil {
		snapshotRestoreService.IgnoreUnavailable(*snapshotRestore.IgnoreUnavailable)
	}
	if len(snapshotRestore.IndexList) > 0 {
		snapshotRestoreService.Indices(snapshotRestore.IndexList...)
	}
	if snapshotRestore.Partial != nil {
		snapshotRestoreService.Partial(*snapshotRestore.Partial)
	}
	if snapshotRestore.IncludeGlobalState != nil {
		snapshotRestoreService.IncludeGlobalState(*snapshotRestore.IncludeGlobalState)
	}
	if snapshotRestore.RenamePattern != "" {
		snapshotRestoreService.RenamePattern(snapshotRestore.RenamePattern)
	}
	if snapshotRestore.RenameReplacement != "" {
		snapshotRestoreService.RenameReplacement(snapshotRestore.RenameReplacement)
	}

	res, err := snapshotRestoreService.Do(ctx)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, res)
}

func (this EsBackUpController) SnapshotStatusAction(ctx *gin.Context) {
	snapshotStatus := es.SnapshotStatus{}
	err := ctx.Bind(&snapshotStatus)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(snapshotStatus.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	snapshotRestoreStatus := esClinet.(*es.EsClientV6).Client.SnapshotStatus().Repository(snapshotStatus.RepositoryName).Snapshot(snapshotStatus.SnapshotName)

	res, err := snapshotRestoreStatus.Do(ctx)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, res)
}
