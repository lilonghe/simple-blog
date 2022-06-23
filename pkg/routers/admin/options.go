package adminRouters

import (
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"strings"
)

func GetOptions(c *gin.Context) {
	options, err := models.GetAllOptions()
	if err != nil {
		panic(err)
	}

	utils.GetCommonResponse(c, options)
}

func UpdateOptions(c *gin.Context) {
	options := make(map[string]string, 0)
	err := c.ShouldBindJSON(&options)
	if err != nil {
		panic(err)
	}

	optionsArr := make([]models.Options, 0)
	for k, v := range options {
		if k == "site_url" && strings.HasSuffix(v, "/") {
			v = strings.TrimRight(v, "/")
		}
		optionsArr = append(optionsArr, models.Options{
			Key:   k,
			Value: v,
		})
	}

	models.UpdateOptions(optionsArr)
	utils.GetCommonSuccess(c)
}
