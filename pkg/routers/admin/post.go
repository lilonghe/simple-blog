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
	keyword := c.Query("keyword")

	condiBean := models.Post{}
	if c.GetString("status") != "" {
		condiBean.Status = int32(c.GetInt("status"))
	}
	if c.GetString("is_public") != "" {
		condiBean.IsPublic = c.GetBool("is_public")
	}

	list, total := models.GetAdminPostList(pageSize, page*pageSize-pageSize, condiBean, keyword)
	utils.GetPageResponse(c, list, total)
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
	}
	utils.GetCommonResponse(c, post)
}
