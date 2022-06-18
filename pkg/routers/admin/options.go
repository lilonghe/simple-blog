package adminRouters

import (
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
)

func GetOptions(c *gin.Context) {
	options, err := models.GetAllOptions()
	if err != nil {
		panic(err)
	}

	utils.GetCommonResponse(c, options)
}
