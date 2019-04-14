package models

import (
	"fmt"
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

	Cates []Cate `json:"cates,omitempty" sql:"-" xorm:"-"`
	Tags  []Tag  `json:"tags,omitempty" sql:"-" xorm:"-"`
}

func GetPostList(limit, offset int) ([]Post, int64, error) {
	datas := make([]Post, 0)
	total, err := global.Store.Count(&Post{})
	err = global.Store.OrderBy("create_time desc").Limit(limit, offset).Find(&datas)
	return datas, total, err
}

func GetPostByPathname(pathname string) (*Post, error) {
	var post Post
	has, err := global.Store.Where(" pathname = ? ", pathname).Get(&post)
	if has {
		return &post, err
	} else {
		return nil, err
	}
}

func GetPostCateAndTag(postId int32) ([]Cate, []Tag) {
	cates := make([]Cate, 0)
	tags := make([]Tag, 0)
	err1 := global.Store.Join("left", PostCate{}.TableName(), fmt.Sprintf("%v.post_id = %d", PostCate{}.TableName(), postId)).Find(&cates)
	err2 := global.Store.Join("left", PostTag{}.TableName(), fmt.Sprintf("%v.post_id = %d", PostTag{}.TableName(), postId)).Find(&tags)
	if err1 != nil {
		fmt.Println("GetPostCateAndTag -> ", err1)
	}
	if err2 != nil {
		fmt.Println("GetPostCateAndTag -> ", err2)
	}

	return cates, tags
}
