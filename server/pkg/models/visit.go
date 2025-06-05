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

func GetWeekVisitListTop(limit int) []map[string]interface{} {
	datas := make([]map[string]interface{}, 0)

	err := global.Store.SQL(`SELECT unique_visits.pathname, title, COUNT(1) AS count
	FROM (SELECT DISTINCT ip, pathname FROM `+Visit{}.TableName()+`
		WHERE create_time >= ?	
	) AS unique_visits
	left join `+Post{}.TableName()+` on `+Post{}.TableName()+`.pathname = unique_visits.pathname
	GROUP BY pathname, title
	ORDER BY count desc
	LIMIT ? `, time.Now().AddDate(0, 0, -7).Format("2006-01-02"), limit).Find(&datas)

	if err != nil {
		panic(err)
	}
	return datas
}

func GetVisitCountByPathList(pathList []string) []map[string]interface{} {
	datas := make([]map[string]interface{}, 0)

	sql := `SELECT unique_visits.pathname, COUNT(1) AS count
	FROM (SELECT DISTINCT ip, pathname FROM ` + Visit{}.TableName() + `) AS unique_visits
	left join ` + Post{}.TableName() + ` on ` + Post{}.TableName() + `.pathname = unique_visits.pathname
	GROUP BY unique_visits.pathname`

	err := global.Store.SQL(sql).In("unique_visits.pathname", pathList).Find(&datas)

	if err != nil {
		panic(err)
	}
	return datas
}
