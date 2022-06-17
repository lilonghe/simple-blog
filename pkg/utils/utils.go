package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

func GetCommonError(err error, c *gin.Context) {
	fmt.Println(err)
	c.JSON(200, map[string]string{"code": "10000", "msg": "Something was wrong"})
}

func GetMessageError(code string, message string, c *gin.Context) {
	c.JSON(200, map[string]string{"code": code, "msg": message})
}

func GetCommonSuccess(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"code": "", "msg": ""})
}

func GetCommonResponse(c *gin.Context, data interface{}) {
	c.JSON(200, map[string]interface{}{"code": "", "msg": "", "data": data})
}
