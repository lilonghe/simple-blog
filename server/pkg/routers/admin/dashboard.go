package adminRouters

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
)

type DashboardResponse struct {
	Count    map[string]interface{} `json:"count"`
	Versions map[string]interface{} `json:"versions"`
}

func Dashboard(c *gin.Context) {
	resp := DashboardResponse{}
	resp.Count = map[string]interface{}{
		"posts": models.GetPostCount(),
		"cates": models.GetCateCount(),
		"tags":  models.GetTagCount(),
	}
	resp.Versions = map[string]interface{}{
		"platform":     runtime.GOOS,
		"mysqlVersion": models.GetDBVersion(),
	}

	utils.GetCommonResponse(c, resp)
}
