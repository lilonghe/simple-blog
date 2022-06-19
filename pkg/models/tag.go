package models

import "simple-blog/pkg/global"

type Tag struct {
	Id       int32  `json:"id,omitempty" xorm:"autoincr"`
	Name     string `json:"name,omitempty"`
	Pathname string `json:"pathname,omitempty"`
}

func (Tag) TableName() string { return global.Config.DbTablePrefix + "tag" }

func GetAllTag() ([]Tag, error) {
	datas := make([]Tag, 0)
	err := global.Store.Find(&datas)
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
