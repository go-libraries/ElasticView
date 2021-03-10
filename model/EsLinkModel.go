package model

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
