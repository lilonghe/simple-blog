package models

import (
	"encoding/json"
	"fmt"
	"html/template"
	"sync"
	"time"

	"simple-blog/pkg/global"
)

type Post struct {
	Id              int32         `json:"id"`
	Status          int32         `json:"status,omitempty"`
	Title           string        `json:"title,omitempty"`
	Pathname        string        `json:"pathname,omitempty"`
	Summary         string        `json:"summary,omitempty"`
	MarkdownContent string        `json:"markdown_content,omitempty"`
	Content         template.HTML `json:"content,omitempty"`
	Options         string        `json:"options,omitempty"`
	CreateTime      time.Time     `json:"create_time,omitempty"`
	UpdateTime      time.Time     `json:"update_time,omitempty"`
	IsPublic        bool          `json:"is_public,omitempty"`

	Cates       []Cate                 `json:"cates,omitempty" sql:"-" xorm:"-"`
	Tags        []Tag                  `json:"tags,omitempty" sql:"-" xorm:"-"`
	PostOptions map[string]interface{} `json:"post_options,omitempty" sql:"-" xorm:"-"`
}

func (Post) TableName() string { return global.Config.DbTablePrefix + "post" }

type PostArchiveView struct {
	Title      string    `json:"title,omitempty"`
	Pathname   string    `json:"pathname,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

func GetPostList(limit, offset int) ([]Post, int64, error) {
	datas := make([]Post, 0)
	total, err := global.Store.Count(&Post{})
	err = global.Store.Where(" status = 3 and is_public = true and type = 0 ").OrderBy("create_time desc").Limit(limit, offset).Find(&datas)
	if err == nil {
		for k, v := range datas {
			err = json.Unmarshal([]byte(v.Options), &datas[k].PostOptions)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return datas, total, err
}

func GetAllPostList(limit, offset int) ([]Post, int64, error) {
	datas := make([]Post, 0)
	total, err := global.Store.Count(&Post{})
	err = global.Store.Where(" status != 4 and is_public = true and type = 0").OrderBy("create_time desc").Limit(limit, offset).Find(&datas)
	return datas, total, err
}

/**
 * Custom action, Only show target cate
 */
func GetPostListByCate(limit, offset int, CateId int32) ([]Post, int64, error) {
	datas := make([]Post, 0)
	total, err := global.Store.Count(&Post{})
	err = global.Store.Where(" status = 3 and is_public = true and id in (select post_id from "+PostCate{}.TableName()+" where cate_id = ? ) ", CateId).OrderBy("create_time desc").Limit(limit, offset).Find(&datas)
	return datas, total, err
}

/**
 * postType: 0-Post, 1-Page
 */
func GetPostByPathname(pathname string, postType int) (*Post, error) {
	var post Post
	has, err := global.Store.Where(" status = 3 and pathname = ? and type = ? ", pathname, postType).Get(&post)
	if has {
		err = json.Unmarshal([]byte(post.Options), &post.PostOptions)
		if err != nil {
			fmt.Println(err)
		}
		return &post, err
	} else {
		return nil, err
	}
}

func GetPostCateAndTag(postId int32) ([]Cate, []Tag) {
	cates := make([]Cate, 0)
	tags := make([]Tag, 0)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		err1 := global.Store.Join("inner", PostCate{}.TableName(), fmt.Sprintf("%v.post_id = %d and %v.cate_id = %v.id", PostCate{}.TableName(), postId, PostCate{}.TableName(), Cate{}.TableName())).Find(&cates)
		if err1 != nil {
			fmt.Println("GetPostCateAndTag -> ", err1)
		}
		wg.Done()
	}()
	go func() {
		err2 := global.Store.Join("inner", PostTag{}.TableName(), fmt.Sprintf("%v.post_id = %d and %v.tag_id = %v.id", PostTag{}.TableName(), postId, PostTag{}.TableName(), Tag{}.TableName())).Find(&tags)
		if err2 != nil {
			fmt.Println("GetPostCateAndTag -> ", err2)
		}
		wg.Done()
	}()
	wg.Wait()
	return cates, tags
}

func GetPostCount() int64 {
	count, err := global.Store.Where("type = 0").Count(Post{})
	if err != nil {
		panic(err)
	}
	return count
}

func GetAdminPostList(limit, offset int, condiBean Post) ([]Post, int64) {
	datas := make([]Post, 0)
	total, err := global.Store.Count(&Post{})
	if err != nil {
		panic(err)
	}
	err = global.Store.Where(" type = 0").OrderBy("create_time desc").Limit(limit, offset).Find(&datas, condiBean)
	if err != nil {
		panic(err)
	}
	return datas, total
}
