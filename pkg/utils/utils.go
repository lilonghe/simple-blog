package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"simple-blog/pkg/global"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	c.JSON(200, map[string]string{"code": "10000", "msg": TranslateMessage(c, "Something was wrong")})
}

func GetMessageError(code string, message string, c *gin.Context) {
	c.JSON(200, map[string]string{"code": code, "msg": TranslateMessage(c, message)})
}

func GetCommonSuccess(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"code": "", "msg": ""})
}

func GetCommonResponse(c *gin.Context, data interface{}) {
	c.JSON(200, map[string]interface{}{"code": "", "msg": "", "data": data})
}

func GetPageResponse(c *gin.Context, list interface{}, total int64) {
	c.JSON(200, map[string]interface{}{"code": "", "msg": "", "data": map[string]interface{}{
		"list":  list,
		"total": total,
	}})
}

func AutoCreateFolder(basePath string, folder string) {
	_, err := os.Stat(path.Join(basePath, folder))
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path.Join(basePath, folder), 0711)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

}

func TranslateMessage(c *gin.Context, key string) string {
	lang := strings.Split(c.GetHeader("Accept-Language"), ",")[0]
	if lang == "" {
		return key
	}
	if global.Translate[key] == nil {
		return key
	}
	if global.Translate[key][lang] == "" {
		return key
	}
	return global.Translate[key][lang]
}
