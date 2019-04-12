package models

import "testing"

func Test_GetAllCate(t *testing.T) {
	_, err := GetAllCate()
	if err != nil {
		t.Error(err)
	}
}
