package models

import "testing"

func Test_GetPostByPathname(t *testing.T) {
	_, err := GetPostByPathname("responsive-f1onts", 1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPostCateAndTag(t *testing.T) {
	GetPostCateAndTag(1)
}
