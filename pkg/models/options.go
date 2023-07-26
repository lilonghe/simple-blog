package models

import (
	"encoding/json"
	"simple-blog/pkg/global"
	"simple-blog/pkg/utils"
	"strconv"
	"time"
)

type Options struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

func GetAllOptions() ([]Options, error) {
	datas := make([]Options, 0)
	err := global.Store.Find(&datas)
	return datas, err
}

func LoadOption() {
	optionArr := make([]Options, 0)
	err := global.Store.Find(&optionArr)
	if err != nil {
		panic(err)
	}

	for _, v := range optionArr {
		global.Options[v.Key] = v.Value
		if v.Key == "themeConfig" {
			err = json.Unmarshal([]byte(v.Value), &global.ThemeConfig)
			if err != nil {
				panic(err)
			}
		} else if v.Key == "comment" {
			if v.Value != "" {
				err = json.Unmarshal([]byte(v.Value), &global.CommentConfig)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	global.Options["author"] = global.Config.Author
	global.Options["authorlink"] = global.Config.AuthorLink

	if global.Config.SiteBegin != nil {
		global.Options["runday"] = strconv.FormatInt(utils.DayDiff(time.Now(), *global.Config.SiteBegin), 10)
	}
}

func GetDBVersion() string {
	r, err := global.Store.QueryString("SELECT VERSION() as version")
	if err != nil {
		panic(err)
	}
	return r[0]["version"]
}

func UpdateOptions(list []Options) {
	sess := global.Store.NewSession()
	for _, v := range list {
		_, err := sess.Update(&v, Options{Key: v.Key})
		if err != nil {
			sess.Rollback()
			panic(err)
		}
	}
}
