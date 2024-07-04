package global

import (
	"fmt"
	"html/template"
	"os"
	"time"

	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/microcosm-cc/bluemonday"
	"xorm.io/xorm"
)

type TranslateMessage map[string]string

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
		CDNHost          string     `yaml:"cdnHost"`
		Host             string     `yaml:"host"`
		EnableAdmin      bool       `yaml:"enableAdmin"`
	}{}

	Options       = map[string]string{}
	ThemeConfig   = map[string]template.HTML{}
	CommentConfig = map[string]template.HTML{}

	HTMLFormat = bluemonday.StrictPolicy()

	Translate = map[string]TranslateMessage{}
)

func init() {
	err := configor.Load(&Config, "config.yml")
	if err != nil {
		panic(err)
	}

	if Config.DbUrl == "" {
		Config.DbUrl = os.Getenv("dbUrl")
	}

	fmt.Println(CommentConfig["name"])
	initDB()
	initTranslate()
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

func initTranslate() {
	configor.Load(&Translate, "i18n.yml")
}
