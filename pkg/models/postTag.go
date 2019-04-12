package models

type PostTag struct {
	Id     int32 `json:"id,omitempty"`
	PostId int32 `json:"post_id,omitempty"`
	TagId  int32 `json:"tag_id,omitempty"`
}
