package models

import "simple-blog/pkg/global"

type Tag struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Pathname string `json:"pathname,omitempty"`
}

func (Tag) TableName() string { return global.Config.DbTablePrefix + "tag" }

func GetAllTag() ([]Tag, error) {
	datas := make([]Tag, 0)
	err := global.Store.Find(&datas)
	return datas, err
}
