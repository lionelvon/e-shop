package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func InitMySQL() {
	engine, _ = xorm.NewEngine("mysql", "eshop:eshop@tcp(localhost:3306)/eshop?charset=utf8")
	engine.DB().SetMaxOpenConns(100)
	engine.ShowSQL(true)
}

func MySQL() *xorm.Engine {
	return engine
}
