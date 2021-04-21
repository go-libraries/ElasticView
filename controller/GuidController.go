package controller

import (
	"github.com/1340691923/ElasticView/engine/db"
	"github.com/1340691923/ElasticView/model"
	"github.com/1340691923/ElasticView/platform-basic-libs/jwt"
	"github.com/1340691923/ElasticView/platform-basic-libs/response"
	"github.com/1340691923/ElasticView/platform-basic-libs/util"

	"github.com/gin-gonic/gin"
)

// 引导控制器
type GuidController struct {
	BaseController
}

// 完成新手引导
func (this GuidController) Finish(ctx *gin.Context) {
	c, err := jwt.ParseToken(ctx.GetHeader("X-Token"))
	if err != nil {
		this.Error(ctx, err)
		return
	}

	gmGuidModel := model.GmGuidModel{}
	err = ctx.Bind(&gmGuidModel)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	_, err = db.SqlBuilder.
		Insert(gmGuidModel.TableName()).
		SetMap(map[string]interface{}{
			"uid":       c.ID,
			"guid_name": gmGuidModel.GuidName,
		}).RunWith(db.Sqlx).Exec()

	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

//是否完成新手引导
func (this GuidController) IsFinish(ctx *gin.Context) {
	c, err := jwt.ParseToken(ctx.GetHeader("X-Token"))
	if err != nil {
		this.Error(ctx, err)
		return
	}

	gmGuidModel := model.GmGuidModel{}
	err = ctx.Bind(&gmGuidModel)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	sql, args, err := db.SqlBuilder.
		Select("count(*)").
		From(gmGuidModel.TableName()).
		Where(db.Eq{
			"uid":       c.ID,
			"guid_name": gmGuidModel.GuidName,
		}).ToSql()

	if err != nil {
		this.Error(ctx, err)
		return
	}
	var count int
	err = db.Sqlx.Get(&count, sql, args...)
	if util.FilterMysqlNilErr(err) {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, count > 0)
}
