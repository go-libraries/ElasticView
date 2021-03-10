package db

import (
	"fmt"
	"log"
	"time"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var Sqlx *sqlx.DB
var SqlBuilder = squirrel.StatementBuilder
type Eq = squirrel.Eq
type Or = squirrel.Or
type Like = squirrel.Like
type SelectBuilder = squirrel.SelectBuilder
type InsertBuilder = squirrel.InsertBuilder
type UpdateBuilder = squirrel.UpdateBuilder

// NewMySQL 创建一个连接MySQL的实体池
func NewSQLX(dbSource string, maxOpenConns, maxIdleConns int){

	db, err := sqlx.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	if maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	}

	if maxIdleConns > 0 {
		db.SetMaxIdleConns(maxIdleConns)
	}

	go func() {
		for {
			err := db.Ping()
			if err != nil {
				log.Println("mysql db can't connect!")
			}
			time.Sleep(time.Minute)
		}
	}()
	Sqlx = db
	return
}

func CreatePage(page, limit int) uint64 {
	tmp := (page - 1) * limit
	return uint64(tmp)
}

func CreateLike(column string) string {
	return fmt.Sprint("%", column, "%")
}
