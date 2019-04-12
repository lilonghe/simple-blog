package models

import (
	"time"
)

type Post struct {
	Id              int32     `json:"id,omitempty"`
	Status          int32     `json:"status,omitempty"`
	Title           string    `json:"title,omitempty"`
	Pathname        string    `json:"pathname,omitempty"`
	Summary         string    `json:"summary,omitempty"`
	MarkdownContent string    `json:"markdown_content,omitempty"`
	Content         string    `json:"content,omitempty"`
	IsPublic        bool      `json:"is_public,omitempty"`
	Options         string    `json:"options,omitempty"`
	CreateTime      time.Time `json:"create_time,omitempty"`
	UpdateTime      time.Time `json:"update_time,omitempty"`
}
