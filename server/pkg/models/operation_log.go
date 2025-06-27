package models

import (
	"time"
	"simple-blog/pkg/global"
)

type OperationLog struct {
	Id            int32     `json:"id"`
	OpType        string    `json:"op_type"`
	OpTargetId    int32     `json:"op_target_id"`
	OpTime        time.Time `json:"op_time"`
	BeforeContent string    `json:"before_content"`
	AfterContent  string    `json:"after_content"`
	OpDetail      string    `json:"op_detail"`
}

func (OperationLog) TableName() string {
	return global.Config.DbTablePrefix + "operation_log"
} 