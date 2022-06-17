package models

import (
	"simple-blog/pkg/global"
)

type Cate struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Pid      int32  `json:"pid,omitempty"`
	Pathname string `json:"pathname,omitempty"`
}

func (Cate) TableName() string { return global.Config.DbTablePrefix + "cate" }

func GetAllCate() ([]Cate, error) {
	datas := make([]Cate, 0)
	err := global.Store.Find(&datas)
	return datas, err
}

func GetCateCount() int64 {
	count, err := global.Store.Count(&Cate{})
	if err != nil {
		panic(err)
	}
	return count
}
