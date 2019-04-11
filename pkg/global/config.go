package global

import (
	"github.com/go-xorm/xorm"
)

var (
	Store *xorm.Engine
)

func init() {
	initDB()
}

func initDB() {
	Store, err := xorm.NewEngine("mysql", "")
	if err != nil {
		panic(err)
	}
	err = Store.Ping()
	if err != nil {
		panic(err)
	}
}
