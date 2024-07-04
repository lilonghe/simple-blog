package models

import (
	"simple-blog/pkg/global"
	"time"
)

type Comment struct {
	Id         int32     `json:"id"`
	Site       string    `json:"site"`
	Email      string    `json:"-"`
	Nickname   string    `json:"nickname"`
	Msg        string    `json:"msg"`
	UserAgent  string    `json:"user_agent"`
	Ip         string    `json:"ip"`
	Audit      bool      `json:"-"`
	CreateTime time.Time `json:"create_time"`
}

type AddCommentViewModel struct {
	Site     string `json:"site"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Msg      string `json:"msg"`
}

func (Comment) TableName() string { return global.Config.DbTablePrefix + "comment" }

func AddComment(c Comment) error {
	d, _ := time.ParseDuration("-24h")
	oneDayAgo := time.Now().Add(d)
	count, err := global.Store.Where("create_time > ? ", oneDayAgo).FindAndCount(&Comment{}, &Comment{Ip: c.Ip})
	if count < 5 {
		_, err = global.Store.InsertOne(c)
		return err
	}
	return err
}

func GetAllAuditComment() ([]Comment, error) {
	list := make([]Comment, 0)
	err := global.Store.Where("audit = 1").OrderBy("create_time desc").Find(&list)
	return list, err
}
