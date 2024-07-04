package models

import (
	"fmt"
	"simple-blog/pkg/global"
)

type Cate struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Pid      int32  `json:"pid,omitempty"`
	Pathname string `json:"pathname,omitempty"`
}

type CateWithCountModel struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Pid      int32  `json:"pid,omitempty"`
	Pathname string `json:"pathname,omitempty"`

	PostCount int `json:"post_count"`
}

func (Cate) TableName() string               { return global.Config.DbTablePrefix + "cate" }
func (CateWithCountModel) TableName() string { return global.Config.DbTablePrefix + "cate" }

func GetAllCate() ([]Cate, error) {
	datas := make([]Cate, 0)
	err := global.Store.Find(&datas)
	return datas, err
}

func GetAllCateWithCount() ([]CateWithCountModel, error) {
	datas := make([]CateWithCountModel, 0)
	cateTable := Cate{}.TableName()
	postCateTable := PostCate{}.TableName()
	postTable := Post{}.TableName()
	err := global.Store.SQL(fmt.Sprintf(`
		select %[1]s.*, (
			select count(1) 
			from %[2]s 
			left join %[3]s on %[3]s.id = %[2]s.post_id
			where cate_id = %[1]s.id and %[3]s.status != 4
		) as post_count 
		from %[1]s
		group by %[1]s.id
	`, cateTable, postCateTable, postTable)).Find(&datas)
	return datas, err
}

func GetCateCount() int64 {
	count, err := global.Store.Count(&Cate{})
	if err != nil {
		panic(err)
	}
	return count
}

func CreateOrEditCate(c Cate) {
	if c.Id != 0 {
		_, err := global.Store.Update(&c, Cate{Id: c.Id})
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

func CheckCateNameOrPathRepeat(name string, pathname string, id int32) bool {
	cate := Cate{}
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
	has, err := sess.Get(&cate)
	if err != nil {
		panic(err)
	}
	return has
}

func DeleteCate(id int32) {
	sess := global.Store.NewSession()
	sess.Begin()
	_, err := sess.Delete(PostCate{CateId: id})
	if err != nil {
		sess.Rollback()
		panic(err)
	}
	_, err = sess.Delete(Cate{Id: id})
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

func GetCate(id int32) *Cate {
	cate := Cate{Id: id}
	has, err := global.Store.Get(&cate)
	if err != nil {
		panic(err)
	}
	if !has {
		return nil
	}
	return &cate
}
