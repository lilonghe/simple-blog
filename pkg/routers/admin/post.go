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

func CreatePost(c *gin.Context) {
	var m models.CreatePostModel
	err := c.ShouldBindJSON(&m)
	if err != nil {
		panic(err)
	}

	existPost, err := models.GetPostByPathname(m.Pathname, m.Type)
	if err != nil {
		panic(err)
	}
	if existPost != nil {
		utils.GetMessageError("REPEAT_PATHNAME", "Repeat pathname", c)
		return
	}

	m.UserId = 1
	models.CreatePost(&m)
	utils.GetCommonResponse(c, m)
}
