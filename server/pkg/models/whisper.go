package models

import (
	"time"
	"simple-blog/pkg/global"
)

type Whisper struct {
	Id         int32     `json:"id" xorm:"pk autoincr"`
	Content    string    `json:"content"`
	IsPublic   bool      `json:"is_public"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (Whisper) TableName() string {
	return global.Config.DbTablePrefix + "whisper"
}

// 分页获取公开 whisper
func GetWhisperList(limit, offset int) ([]Whisper, int64, error) {
	whispers := make([]Whisper, 0)
	total, err := global.Store.Where("is_public = ?", true).Count(&Whisper{})
	if err != nil {
		return nil, 0, err
	}
	err = global.Store.Where("is_public = ?", true).OrderBy("create_time desc").Limit(limit, offset).Find(&whispers)
	return whispers, total, err
}

// 管理端分页获取 whisper（支持公开性筛选）
func GetWhisperListAdmin(limit, offset int, isPublic *bool) ([]Whisper, int64, error) {
	whispers := make([]Whisper, 0)
	sess := global.Store.NewSession().OrderBy("create_time desc")
	if isPublic != nil {
		sess = sess.Where("is_public = ?", *isPublic)
	}
	total, err := sess.Limit(limit, offset).FindAndCount(&whispers)
	return whispers, total, err
}

// 新建 whisper
func CreateWhisper(w *Whisper) error {
	_, err := global.Store.Insert(w)
	return err
}

// 根据 id 获取 whisper
func GetWhisperById(id int32) (*Whisper, error) {
	var w Whisper
	has, err := global.Store.ID(id).Get(&w)
	if err != nil || !has {
		return nil, err
	}
	return &w, nil
}

// 编辑 whisper
func UpdateWhisper(id int32, w *Whisper) error {
	_, err := global.Store.ID(id).UseBool("is_public").Update(w)
	return err
}

// 删除 whisper
func DeleteWhisper(id int32) error {
	_, err := global.Store.ID(id).Delete(&Whisper{})
	return err
}

// 只更新公开性
func UpdateWhisperVisibility(id int32, isPublic bool) error {
	_, err := global.Store.ID(id).Cols("is_public").Update(&Whisper{IsPublic: isPublic})
	return err
} 