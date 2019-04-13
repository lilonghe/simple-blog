package models

import (
	"html/template"
	"time"

	"github.com/lilonghe/simple-blog/pkg/global"
)

type Post struct {
	Id              int32         `json:"id,omitempty"`
	Status          int32         `json:"status,omitempty"`
	Title           string        `json:"title,omitempty"`
	Pathname        string        `json:"pathname,omitempty"`
	Summary         template.HTML `json:"summary,omitempty"`
	MarkdownContent string        `json:"markdown_content,omitempty"`
	Content         template.HTML `json:"content,omitempty"`
	IsPublic        bool          `json:"is_public,omitempty"`
	Options         string        `json:"options,omitempty"`
	CreateTime      time.Time     `json:"create_time,omitempty"`
	UpdateTime      time.Time     `json:"update_time,omitempty"`
}

func GetPostList(limit, offset int) ([]Post, int64, error) {
	datas := make([]Post, 0)
	total, err := global.Store.Count(&Post{})
	err = global.Store.OrderBy("create_time desc").Limit(limit, offset).Find(&datas)
	return datas, total, err
}
