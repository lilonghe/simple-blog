package adminRouters

import (
	"runtime"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

type DashboardResponse struct {
	Count    map[string]interface{} `json:"count"`
	Versions map[string]interface{} `json:"versions"`
}

func Dashboard(c *gin.Context) {
	resp := DashboardResponse{}
	resp.Count = map[string]interface{}{
		"posts":        models.GetPostCount(),
		"cates":        models.GetCateCount(),
		"tags":         models.GetTagCount(),
		"weekVisitTop": models.GetWeekVisitListTop(5),
	}
	resp.Versions = map[string]interface{}{
		"platform":     runtime.GOOS,
		"mysqlVersion": models.GetDBVersion(),
	}

	utils.GetCommonResponse(c, resp)
}
