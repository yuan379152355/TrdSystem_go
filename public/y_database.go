package tspublic

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库句柄
type SYDatabase struct {
	Db *sqlx.DB
}

func (sDb *SYDatabase) Init() bool {
	db, err := sqlx.Open("mysql", "root:1@tcp(127.0.0.1:3306)/microservice?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return false
	}
	sDb.Db = db
	return true
}
