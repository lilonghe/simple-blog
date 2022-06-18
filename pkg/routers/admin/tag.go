package adminRouters

import (
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
)

func GetTagList(c *gin.Context) {
	list, err := models.GetAllTag()
	if err != nil {
		panic(err)
	}

	utils.GetCommonResponse(c, list)
}
