package global

import (
	"fmt"
	"github.com/lilonghe/simple-blog/pkg/utils"
	"html/template"

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
		DbUrl         string `default:"root:root@/blog?charset=utf8"`
		DbTablePrefix string `default:"fk_"`
	}{}

	Options     = map[string]string{}
	ThemeConfig = map[string]template.HTML{}
)

func init() {
	configor.Load(&Config, "config.yml")
	fmt.Println("Config -> ", utils.ToJson(Config))
	initDB()
}

func initDB() {
	engine, err := xorm.NewEngine("mysql", Config.DbUrl)
	if err != nil {
		panic(err)
	}
	Store = engine

	if Config.Dev {
		Store.ShowSQL(true)
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, Config.DbTablePrefix)
	Store.SetTableMapper(tbMapper)
}
