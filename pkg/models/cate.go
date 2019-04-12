package models

import (
	"time"
)

type Cate struct {
	Id         int32     `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Pid        int32     `json:"pid,omitempty"`
	Pathname   string    `json:"pathname,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}
