package models

import "github.com/lilonghe/simple-blog/pkg/global"

type PostCate struct {
	Id     int32 `json:"id,omitempty"`
	PostId int32 `json:"post_id,omitempty"`
	CateId int32 `json:"cate_id,omitempty"`
}

func (PostCate) TableName() string { return global.Config.DbTablePrefix + "post_cate" }
