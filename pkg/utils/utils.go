package utils

import "encoding/json"

func ToJson(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}
