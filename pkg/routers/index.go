package routers

import (
	"fmt"
	"github.com/lilonghe/simple-blog/pkg/global"
	"strconv"
	"strings"

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
		"options":     global.Options,
		"themeConfig": global.ThemeConfig,

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

func PostDetail(c *gin.Context) {
	paths := strings.Split(c.Param("pathname"), ".html")
	pathname := ""
	if len(paths) > 0 {
		pathname = paths[0]
	}

	post, _ := models.GetPostByPathname(pathname)
	resp := map[string]interface{}{
		"options":     global.Options,
		"themeConfig": global.ThemeConfig,
	}

	fmt.Println(global.ThemeConfig)

	if post != nil {
		cates, tags := models.GetPostCateAndTag(post.Id)
		post.Cates = cates
		post.Tags = tags
		resp["post"] = *post
	} else {
		c.Redirect(301, "/")
	}

	c.HTML(200, "post.html", resp)
}
