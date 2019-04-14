package models

import (
	"encoding/json"
	"github.com/lilonghe/simple-blog/pkg/global"
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
		}
	}
}