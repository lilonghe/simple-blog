package models

import (
	"simple-blog/pkg/global"
	"time"
)

type Visit struct {
	Id         int32     `json:"id" xorm:"pk autoincr"`
	Pathname   string    `json:"pathname"`
	UserAgent  string    `json:"user_agent"`
	Ip         string    `json:"ip"`
	CreateTime time.Time `json:"create_time"`
}

func (Visit) TableName() string { return global.Config.DbTablePrefix + "visit" }

func AddVisit(v Visit) error {
	_, err := global.Store.InsertOne(v)
	return err
}

func GetVisitList(limit, offset int, keyword string) ([]Visit, int64) {
	datas := make([]Visit, 0)
	sess := global.Store.Where("1 = 1")
	if keyword != "" {
		sess = sess.Where("user_agent like ?", "%"+keyword+"%")
	}

	total, err := sess.OrderBy("create_time desc").Limit(limit, offset).FindAndCount(&datas)
	if err != nil {
		panic(err)
	}
	return datas, total
}
