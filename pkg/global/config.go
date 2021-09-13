package global

import (
	"html/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/configor"
	"github.com/microcosm-cc/bluemonday"
)

var (
	Store *xorm.Engine
)

var (
	Config = struct {
		Dev           bool
		DbUrl         string `default:"root:root@/blog?charset=utf8"`
		DbTablePrefix string `default:"fk_"`
		Author        string
		AuthorLink    string
		SiteBegin     *time.Time
	}{}

	Options     = map[string]string{}
	ThemeConfig = map[string]template.HTML{}

	HTMLFormat = bluemonday.StripTagsPolicy()
)

func init() {
	configor.Load(&Config, "config.yml")
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
