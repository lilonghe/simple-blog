package adminRouters

import (
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"strconv"
)

func GetPostList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	page, _ := strconv.Atoi(c.Query("page"))

	condiBean := models.Post{}
	if c.GetString("status") != "" {
		condiBean.Status = int32(c.GetInt("status"))
	}
	if c.GetString("is_public") != "" {
		condiBean.IsPublic = c.GetBool("is_public")
	}

	list, total := models.GetAdminPostList(pageSize, page*pageSize-pageSize, condiBean)
	utils.GetPageResponse(c, list, total)
}
