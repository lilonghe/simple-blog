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
	AllowComment    bool          `json:"allow_comment"`

	Cates       []Cate                 `json:"cates,omitempty" sql:"-" xorm:"-"`
	Tags        []Tag                  `json:"tags,omitempty" sql:"-" xorm:"-"`
	PostOptions map[string]interface{} `json:"post_options,omitempty" sql:"-" xorm:"-"`
	VisitCount  int                    `json:"visit_count,omitempty" sql:"-" xorm:"-"`
}

type CreatePostModel struct {
	Id              int32         `json:"id" xorm:"pk autoincr"`
	Status          int32         `json:"status"`
	Title           string        `json:"title"`
	Pathname        string        `json:"pathname"`
	Summary         string        `json:"summary"`
	MarkdownContent string        `json:"markdown_content"`
	Content         template.HTML `json:"content"`
	Options         *string       `json:"options"`
	CreateTime      time.Time     `json:"create_time"`
	UpdateTime      time.Time     `json:"update_time"`
	IsPublic        bool          `json:"is_public"`
	AllowComment    bool          `json:"allow_comment"`
	UserId          int32         `json:"-"`
	Type            int           `json:"type"`

	CateList []int32  `json:"cate_list" sql:"-" xorm:"-"`
	TagList  []string `json:"tag_list" sql:"-" xorm:"-"`
}

func (Post) TableName() string            { return global.Config.DbTablePrefix + "post" }
func (CreatePostModel) TableName() string { return global.Config.DbTablePrefix + "post" }

type PostArchiveView struct {
	Title      string    `json:"title,omitempty"`
	Pathname   string    `json:"pathname,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

type PostListFilterViewModal struct {
	Status     *int
	IsPublic   *bool
	CreateTime *time.Time
}

func GetPostList(limit, offset int) ([]Post, int64, error) {
	datas := make([]Post, 0)
	total, err := global.Store.Count(&Post{})
	err = global.Store.Where(" status = 3 and is_public = true and type = 0 ").OrderBy("create_time desc").Limit(limit, offset).Find(&datas)
	if err == nil {
		for k, v := range datas {
			if v.Options != "" {
				err = json.Unmarshal([]byte(v.Options), &datas[k].PostOptions)
				if err != nil {
					fmt.Println(err)
				}
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
func GetPublishedPostByPathname(pathname string, postType int) (*Post, error) {
	var post Post
	has, err := global.Store.Where(" status = 3 and pathname = ? and type = ? ", pathname, postType).Get(&post)
	if has {
		if post.Options != "" {
			fmt.Println(post.Options)
			err = json.Unmarshal([]byte(post.Options), &post.PostOptions)
			if err != nil {
				fmt.Println(err)
			}
		}
		return &post, err
	} else {
		return nil, err
	}
}

/**
 * postType: 0-Post, 1-Page
 */
func GetPostByPathname(pathname string, postType int) (*Post, error) {
	var post Post
	has, err := global.Store.Where("status != 4 and pathname = ? and type = ? ", pathname, postType).Get(&post)
	if has {
		if post.Options != "" {
			err = json.Unmarshal([]byte(post.Options), &post.PostOptions)
			if err != nil {
				fmt.Println(err)
			}
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
	count, err := global.Store.Where("status != 4 and type = 0").Count(Post{})
	if err != nil {
		panic(err)
	}
	return count
}

func GetAdminPostList(limit, offset int, condiBean PostListFilterViewModal, keyword string, pageType int) ([]Post, int64) {
	datas := make([]Post, 0)
	sess := global.Store.Where(" status != 4 and type = ?", pageType)

	if keyword != "" {
		sess = sess.Where("title like ?", "%"+keyword+"%")
	}
	if condiBean.Status != nil {
		sess = sess.Where("status = ?", *condiBean.Status)
	}
	if condiBean.IsPublic != nil {
		sess = sess.Where("is_public = ?", *condiBean.IsPublic)
	}
	if condiBean.CreateTime != nil {
		sess = sess.Where("create_time >= ?", *condiBean.CreateTime)
	}

	total, err := sess.OrderBy("create_time desc").Limit(limit, offset).FindAndCount(&datas)
	if err != nil {
		panic(err)
	}
	return datas, total
}

func CreateOrEditPost(post *CreatePostModel) {
	sess := global.Store.NewSession()
	sess.Begin()

	if post.Id == 0 {
		_, err := sess.Insert(post)
		if err != nil {
			sess.Rollback()
			panic(err)
		}
	} else {
		/**
		* MustCols("is_public", "allow_comment", "options")
		* condiBean 和 MustCols 冲突，所以改成了 column 加 pk，然后用 ID()
		* 或者使用 Where("id = ? ", id) 也可
		* 但是 options 仍然不更新，没辙只能换成 AllCols
		 */
		_, err := sess.ID(post.Id).AllCols().Update(post)
		if err != nil {
			sess.Rollback()
			panic(err)
		}
	}

	_, err := sess.Delete(PostCate{PostId: post.Id})
	if err != nil {
		sess.Rollback()
		panic(err)
	}

	if len(post.CateList) > 0 {
		postCate := make([]PostCate, 0)
		for _, v := range post.CateList {
			postCate = append(postCate, PostCate{PostId: post.Id, CateId: v})
		}
		_, err = sess.InsertMulti(&postCate)
		if err != nil {
			sess.Rollback()
			panic(err)
		}
	}

	_, err = sess.Delete(PostTag{PostId: post.Id})
	if err != nil {
		sess.Rollback()
		panic(err)
	}

	if len(post.TagList) > 0 {
		postTag := make([]PostTag, 0)
		for _, v := range post.TagList {
			tag := Tag{Name: v}
			has, err := sess.Get(&tag)
			if err != nil {
				sess.Rollback()
				panic(err)
			}
			if !has {
				tag.Pathname = tag.Name
				_, err = sess.Insert(&tag)
				if err != nil {
					sess.Rollback()
					panic(err)
				}
			}
			postTag = append(postTag, PostTag{PostId: post.Id, TagId: tag.Id})
		}
		_, err = sess.InsertMulti(&postTag)
		if err != nil {
			sess.Rollback()
			panic(err)
		}
	}

	err = sess.Commit()
	if err != nil {
		sess.Rollback()
		panic(err)
	}
}

func GetEditPost(id int32) *CreatePostModel {
	p := CreatePostModel{Id: id}
	has, err := global.Store.Get(&p)
	if err != nil {
		panic(err)
	}
	if !has {
		return nil
	}

	p.CateList = make([]int32, 0)
	p.TagList = make([]string, 0)
	cates, tags := GetPostCateAndTag(id)
	for _, v := range cates {
		p.CateList = append(p.CateList, v.Id)
	}
	for _, v := range tags {
		p.TagList = append(p.TagList, v.Name)
	}

	return &p
}

func DeletePost(id int32) {
	_, err := global.Store.Update(Post{Id: id, Status: 4, UpdateTime: time.Now()}, Post{Id: id})
	if err != nil {
		panic(err)
	}
}
