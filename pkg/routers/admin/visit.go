package adminRouters

import (
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetVisitList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	page, _ := strconv.Atoi(c.Query("page"))
	keyword := c.Query("keyword")

	list, total := models.GetVisitList(pageSize, page*pageSize-pageSize, keyword)
	utils.GetPageResponse(c, list, total)
}
