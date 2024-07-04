package models

import "testing"

func Test_GetAllOptions(t *testing.T) {
	_, err := GetAllOptions()
	if err != nil {
		t.Error(err)
	}
}
