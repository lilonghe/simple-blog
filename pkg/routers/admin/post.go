package adminRouters

import (
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	page, _ := strconv.Atoi(c.Query("page"))
	keyword := c.Query("keyword")
	archive := c.Query("archive")

	condiBean := models.Post{}
	if c.GetString("status") != "" {
		condiBean.Status = int32(c.GetInt("status"))
	}
	if c.GetString("is_public") != "" {
		condiBean.IsPublic = c.GetBool("is_public")
	}

	list, total := models.GetAdminPostList(pageSize, page*pageSize-pageSize, condiBean, keyword)

	if archive == "true" {
		archiveList := make([]models.PostArchiveView, 0)
		for _, v := range list {
			archiveList = append(archiveList, models.PostArchiveView{
				Title:      v.Title,
				Pathname:   v.Pathname,
				CreateTime: v.CreateTime,
				UpdateTime: v.UpdateTime,
			})
		}
		utils.GetPageResponse(c, archiveList, total)
	} else {
		utils.GetPageResponse(c, list, total)
	}

}

func CreateOrEditPost(c *gin.Context) {
	var m models.CreatePostModel
	err := c.ShouldBindJSON(&m)
	if err != nil {
		panic(err)
	}

	existPost, err := models.GetPostByPathname(m.Pathname, m.Type)
	if err != nil {
		panic(err)
	}
	if m.Id == 0 {
		if existPost != nil {
			utils.GetMessageError("REPEAT_PATHNAME", "Repeat pathname", c)
			return
		}
	} else {
		if existPost == nil {
			utils.GetMessageError("POST_NOT_FOUND", "Post not found", c)
			return
		}
	}

	m.UserId = int32(c.GetInt("id"))
	models.CreateOrEditPost(&m)
	utils.GetCommonResponse(c, m)
}

func GetEditPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	post := models.GetEditPost(int32(id))
	if post == nil {
		utils.GetMessageError("POST_NOT_FOUND", "Post not found", c)
		return
	}
	utils.GetCommonResponse(c, post)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post := models.GetEditPost(int32(id))
	if post == nil {
		utils.GetMessageError("POST_NOT_FOUND", "Post not found", c)
		return
	}

	models.DeletePost(int32(id))
	utils.GetCommonSuccess(c)
}
