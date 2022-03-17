package models

import (
	"simple-blog/pkg/global"
	"time"
)

type Visit struct {
	Id         int32
	Pathname   string
	UserAgent  string
	Ip         string
	CreateTime time.Time
}

func (Visit) TableName() string { return global.Config.DbTablePrefix + "visit" }

func AddVisit(v Visit) error {
	_, err := global.Store.InsertOne(v)
	return err
}
