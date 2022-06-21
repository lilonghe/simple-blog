package models

import (
	"fmt"
	"simple-blog/pkg/global"
)

type Tag struct {
	Id       int32  `json:"id,omitempty" xorm:"autoincr"`
	Name     string `json:"name,omitempty"`
	Pathname string `json:"pathname,omitempty"`
}

type TagWithCountModel struct {
	Id       int32  `json:"id,omitempty" xorm:"autoincr"`
	Name     string `json:"name,omitempty"`
	Pathname string `json:"pathname,omitempty"`

	PostCount int `json:"post_count"`
}

func (Tag) TableName() string               { return global.Config.DbTablePrefix + "tag" }
func (TagWithCountModel) TableName() string { return global.Config.DbTablePrefix + "tag" }

func GetAllTag() ([]Tag, error) {
	datas := make([]Tag, 0)
	err := global.Store.Find(&datas)
	return datas, err
}

func GetAllTagWithCount() ([]TagWithCountModel, error) {
	datas := make([]TagWithCountModel, 0)
	tagTable := Tag{}.TableName()
	postTagTable := PostTag{}.TableName()
	postTable := Post{}.TableName()
	err := global.Store.SQL(fmt.Sprintf(`
		select %[1]s.*, (
			select count(1) 
			from %[2]s 
			left join %[3]s on %[3]s.id = %[2]s.post_id
			where tag_id = %[1]s.id and %[3]s.status != 4
		) as post_count 
		from %[1]s
		group by %[1]s.id
	`, tagTable, postTagTable, postTable)).Find(&datas)
	return datas, err
}

func GetTagCount() int64 {
	count, err := global.Store.Count(&Tag{})
	if err != nil {
		panic(err)
	}
	return count
}

func CreateOrEditTag(c Tag) {
	if c.Id != 0 {
		_, err := global.Store.Update(&c, Tag{Id: c.Id})
		if err != nil {
			panic(err)
		}
	} else {
		_, err := global.Store.InsertOne(&c)
		if err != nil {
			panic(err)
		}
	}
}

func DeleteTag(id int32) {
	sess := global.Store.NewSession()
	sess.Begin()
	_, err := sess.Delete(PostTag{TagId: id})
	if err != nil {
		sess.Rollback()
		panic(err)
	}
	_, err = sess.Delete(Tag{Id: id})
	if err != nil {
		sess.Rollback()
		panic(err)
	}
	err = sess.Commit()
	if err != nil {
		sess.Rollback()
		panic(err)
	}
}

func GetTag(id int32) *Tag {
	tag := Tag{Id: id}
	has, err := global.Store.Get(&tag)
	if err != nil {
		panic(err)
	}
	if !has {
		return nil
	}
	return &tag
}

func CheckTagNameOrPathRepeat(name string, pathname string, id int32) bool {
	tag := Tag{}
	sess := global.Store.NewSession()
	if name != "" {
		sess = sess.Or("name = ?", name)
	}
	if pathname != "" {
		sess = sess.Or("pathname = ?", pathname)
	}
	if id != 0 {
		sess = sess.Where("id != ?", id)
	}
	has, err := sess.Get(&tag)
	if err != nil {
		panic(err)
	}
	return has
}
