package models

import "testing"

func Test_GetPostByPathname(t *testing.T) {
	_, err := GetPublishedPostByPathname("responsive-f1onts", 1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPostCateAndTag(t *testing.T) {
	GetPostCateAndTag(1)
}

func TestGetAdminPostList(t *testing.T) {
	//list, count, err := GetAdminPostList(10, 0, &Post{
	//	Status: 1,
	//})
	//t.Log(list)
	//t.Log(count)
	//t.Log(err)
}
