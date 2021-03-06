package models

import "simple-blog/pkg/global"

type PostCate struct {
	Id     int32 `json:"id,omitempty"`
	PostId int32 `json:"post_id,omitempty"`
	CateId int32 `json:"cate_id,omitempty"`
}

func (PostCate) TableName() string { return global.Config.DbTablePrefix + "post_cate" }

func GetAllPostCate() ([]PostCate, error) {
	datas := make([]PostCate, 0)
	err := global.Store.Find(&datas)
	return datas, err
}
