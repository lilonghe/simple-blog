package models

type PostCate struct {
	Id     int32 `json:"id,omitempty"`
	PostId int32 `json:"post_id,omitempty"`
	CateId int32 `json:"cate_id,omitempty"`
}
