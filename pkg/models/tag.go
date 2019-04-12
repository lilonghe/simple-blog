package models

type Tag struct {
	Id       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Pathname string `json:"pathname,omitempty"`
}
