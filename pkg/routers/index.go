package routers

import (
	"strconv"
	"strings"

	"github.com/lilonghe/simple-blog/pkg/global"

	"github.com/gin-gonic/gin"
	"github.com/lilonghe/simple-blog/pkg/models"
)

func Index(c *gin.Context) {
	limit, _ := strconv.Atoi(global.Options["postsListSize"])
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

type archivesModel struct {
	Year  string
	Month string
	List  []models.PostArchiveView
}

func Archives(c *gin.Context) {
	posts, _, _ := models.GetPostList(30000, 0)
	respData := make([]archivesModel, 0)

	for _, v := range posts {
		year := strconv.FormatInt(int64(v.CreateTime.Year()), 10)
		month := v.CreateTime.Format("01")
		match := false
		for k, item := range respData {
			if item.Year == year && item.Month == month {
				respData[k].List = append(respData[k].List, models.PostArchiveView{
					Pathname:   v.Pathname,
					Title:      v.Title,
					CreateTime: v.CreateTime,
					UpdateTime: v.UpdateTime,
				})
				match = true
				break
			}
		}
		if !match {
			respData = append(respData, archivesModel{
				Year:  year,
				Month: month,
				List: []models.PostArchiveView{{
					Pathname:   v.Pathname,
					Title:      v.Title,
					CreateTime: v.CreateTime,
					UpdateTime: v.UpdateTime,
				}},
			})
		}
	}

	resp := map[string]interface{}{
		"options": global.Options,
		"datas":   respData,
	}

	c.HTML(200, "archive.html", resp)
}
