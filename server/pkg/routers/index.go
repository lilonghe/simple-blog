package routers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"simple-blog/pkg/global"
	"simple-blog/pkg/utils"

	"simple-blog/pkg/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	limit, _ := strconv.Atoi(global.Options["postsListSize"])
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	posts, total, _ := models.GetPostList(limit, (page-1)*limit)

	for k, v := range posts {
		posts[k].Summary = utils.GetPostSummary(v.Summary, 140)
	}

	resp := map[string]interface{}{
		"options":     utils.GetRenderOptions(),
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
	preview := c.Query("preview")
	pathname := ""
	if len(paths) > 0 {
		pathname = paths[0]
	}

	var post *models.Post
	if preview == "true" {
		post, _ = models.GetPostByPathname(pathname, 0)
	} else {
		post, _ = models.GetPublishedPostByPathname(pathname, 0)
	}
	resp := map[string]interface{}{
		"options":       utils.GetRenderOptions(),
		"themeConfig":   global.ThemeConfig,
		"commentConfig": global.CommentConfig,
	}

	if post != nil {
		cates, tags := models.GetPostCateAndTag(post.Id)
		post.Cates = cates
		post.Tags = tags
		post.Summary = utils.GetPostSummary(post.Summary, 200)
		post.Content = utils.GetUseCDNContent(post.Content)
		resp["post"] = *post
	} else {
		c.Redirect(301, "/")
	}

	visitInfo := models.Visit{
		Pathname:   pathname,
		UserAgent:  c.GetHeader("User-Agent"),
		Ip:         c.GetHeader("X-Real-IP"),
		CreateTime: time.Now(),
	}
	err := models.AddVisit(visitInfo)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(200, "post.html", resp)
}

func PageDetail(c *gin.Context) {
	paths := strings.Split(c.Param("pathname"), ".html")
	pathname := ""
	if len(paths) > 0 {
		pathname = paths[0]
	}

	post, _ := models.GetPublishedPostByPathname(pathname, 1)
	resp := map[string]interface{}{
		"options":       utils.GetRenderOptions(),
		"themeConfig":   global.ThemeConfig,
		"commentConfig": global.CommentConfig,
	}

	if post != nil {
		post.Summary = utils.GetPostSummary(post.Summary, 200)
		post.Content = utils.GetUseCDNContent(post.Content)
		resp["post"] = *post
		resp["pathname"] = pathname
	} else {
		c.Redirect(301, "/")
	}

	c.HTML(200, "page.html", resp)
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
		"options": utils.GetRenderOptions(),
		"datas":   respData,
	}

	c.HTML(200, "archive.html", resp)
}

func SitemapTxt(c *gin.Context) {
	text, _ := url.JoinPath(global.Options["site_url"], "archives")
	postList, _, _ := models.GetPostList(99999, 0)
	for _, v := range postList {
		link, _ := url.JoinPath(global.Options["site_url"], "post", v.Pathname)
		text = text + "\n" + link
	}

	c.String(200, text)
}
