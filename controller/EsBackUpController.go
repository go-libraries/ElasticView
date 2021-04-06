package controller

import (
	"ElasticView/engine/es"
	"ElasticView/engine/logs"
	"ElasticView/platform-basic-libs/response"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type EsBackUpController struct {
	BaseController
}

func (this EsBackUpController) SnapshotListAction(ctx *gin.Context) {
	esSnapshotInfo := es.EsSnapshotInfo{}
	err = ctx.Bind(&esSnapshotInfo)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(esSnapshotInfo.EsConnect)
	if err != nil {
		this.Error(ctx, err)
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
		"list": list,
		"res":  res,
	})
}

func (this EsBackUpController) SnapshotCreateRepositoryAction(ctx *gin.Context) {
	snapshotCreateRepository := es.SnapshotCreateRepository{}
	err = ctx.Bind(&snapshotCreateRepository)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	esClinet, err := es.GetEsClientV6ByID(snapshotCreateRepository.EsConnect)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	settings := map[string]interface{}{
		"location": snapshotCreateRepository.Location,
	}

	if snapshotCreateRepository.Compress != nil {
		compress := *snapshotCreateRepository.Compress
		settings["compress"] = compress
	}

	if snapshotCreateRepository.MaxRestoreBytesPerSec != "" {
		settings["max_restore_bytes_per_sec"] = snapshotCreateRepository.MaxRestoreBytesPerSec
	}

	if snapshotCreateRepository.MaxSnapshotBytesPerSec != "" {
		settings["max_snapshot_bytes_per_sec"] = snapshotCreateRepository.MaxSnapshotBytesPerSec
	}

	_, err = esClinet.(*es.EsClientV6).Client.SnapshotCreateRepository(snapshotCreateRepository.Repository).Type("fs").Settings(
		settings,
	).Do(ctx)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this EsBackUpController) CreateSnapshotAction(ctx *gin.Context) {

}

func (this EsBackUpController) SnapshotDeleteAction(ctx *gin.Context) {

}
