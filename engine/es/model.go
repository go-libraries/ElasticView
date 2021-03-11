package es

import "ElasticView/engine/db"

type Sort struct {
	Field     string
	Ascending bool
}

type Page struct {
	PageNum  int
	PageSize int
}

func NewPage(page int, pageSize int) *Page {
	return &Page{
		PageNum:  int(db.CreatePage(page, pageSize)),
		PageSize: pageSize,
	}
}
