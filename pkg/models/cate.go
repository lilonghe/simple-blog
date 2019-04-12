package models

import (
	"github.com/lilonghe/simple-blog/pkg/global"
)

type Cate struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Pid      int32  `json:"pid,omitempty"`
	Pathname string `json:"pathname,omitempty"`
}

func GetAllCate() ([]Cate, error) {
	datas := make([]Cate, 0)
	err := global.Store.Find(&datas)
	return datas, err
}
