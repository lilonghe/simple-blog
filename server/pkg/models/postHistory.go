package models

import "time"

type PostHistory struct {
	Id              int32     `json:"id,omitempty"`
	PostId          int32     `json:"post_id,omitempty"`
	MarkdownContent string    `json:"markdown_content,omitempty"`
	CreateTime      time.Time `json:"create_time,omitempty"`
}
