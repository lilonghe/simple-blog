package models

import "github.com/lilonghe/simple-blog/pkg/global"

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
