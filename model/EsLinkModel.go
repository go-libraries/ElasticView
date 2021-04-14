package model

import (
	"ElasticView/engine/db"
)

type EsLinkModel struct {
	ID      int64  `gorm:"column:id" json:"id" db:"id"`
	Ip      string `gorm:"column:ip" json:"ip" db:"ip"`
	User    string `gorm:"column:user" json:"user" db:"user"`
	Pwd     string `gorm:"column:pwd" json:"pwd" db:"pwd"`
	Created string `gorm:"column:created" json:"created" db:"created"`
	Updated string `gorm:"column:updated" json:"updated" db:"updated"`
	Remark  string `gorm:"column:remark" json:"remark" db:"remark"`
	Version int    `json:"version" db:"version"`
}

var EsLinkList = []EsLinkModel{}

func (this *EsLinkModel) FlushEsLinkList() (err error) {
	list, err := this.GetListAction()
	if err != nil {
		return
	}
	EsLinkList = list
	return
}

func (this *EsLinkModel) GetListAction() (esLinkList []EsLinkModel, err error) {
	sql, args, err := db.SqlBuilder.
		Select("*").
		From("es_link").ToSql()
	if err != nil {
		return
	}

	err = db.Sqlx.Select(&esLinkList, sql, args...)
	if err != nil {
		return
	}
	return
}
