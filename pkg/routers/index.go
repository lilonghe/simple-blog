package routers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lilonghe/simple-blog/pkg/models"
)

func Index(c *gin.Context) {
	limit := 10
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	posts, total, _ := models.GetPostList(limit, (page-1)*limit)
	resp := map[string]interface{}{
		"posts": posts,
		"total": total,
		"page":  page,
		"limit": limit,
	}

	if page*limit < int(total) {
		resp["next"] = page + 1
	}

	if page > 1 {
		resp["prev"] = page - 1
	}

	c.HTML(200, "index.html", resp)
}
