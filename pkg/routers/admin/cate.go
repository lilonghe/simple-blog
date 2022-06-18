package adminRouters

import (
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
)

func GetCateList(c *gin.Context) {
	list, err := models.GetAllCate()
	if err != nil {
		panic(err)
	}

	utils.GetCommonResponse(c, list)
}
