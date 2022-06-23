package global

import (
	"html/template"
	"time"
	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/microcosm-cc/bluemonday"
	"xorm.io/xorm"
)

var (
	Store *xorm.Engine
)

var (
	Config = struct {
		Dev              bool
		DbUrl            string     `yaml:"dbUrl"`
		DbTablePrefix    string     `yaml:"dbTablePrefix"`
		Author           string     `yaml:"author"`
		AuthorLink       string     `yaml:"authorLink"`
		SiteBegin        *time.Time `yaml:"siteBegin"`
		UploadPath       string     `yaml:"uploadPath"`
		UploadAccessPath string     `yaml:"uploadAccessPath"`
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

	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, Config.DbTablePrefix)
	Store.SetTableMapper(tbMapper)
}
