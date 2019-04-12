package global

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/configor"
)

var (
	Store *xorm.Engine
)

var (
	Config = struct {
		Dev           bool
		DbUrl         string
		DbTablePrefix string
	}{}
)

func init() {
	configor.Load(&Config, "config.yml")
	fmt.Println(Config.Dev)
	initDB()
}

func initDB() {
	engine, err := xorm.NewEngine("mysql", "root:root@/blog?charset=utf8")
	if err != nil {
		panic(err)
	}
	Store = engine
	err = Store.Ping()
	if err != nil {
		panic(err)
	}
	Store.ShowSQL(true)
	Store.SetMaxIdleConns(5)

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "fk_")
	Store.SetTableMapper(tbMapper)
}
