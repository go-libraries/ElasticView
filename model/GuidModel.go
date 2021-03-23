package model

type GmGuidModel struct {
	ID       int    `json:"id"`
	Uid      int    `json:"uid" db:"uid"`
	GuidName string `json:"guid_name" db:"guid_name"`
	Created  string `json:"created" db:"created"`
}

func (this *GmGuidModel) TableName() string {
	return "gm_guid"
}
