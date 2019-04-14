package utils

import (
	"encoding/json"
	"time"
)

func ToJson(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

func DayDiff(t1, t2 time.Time) int64 {
	return abs((t1.Unix() - t2.Unix()) / 86400)
}

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
