package adminRouters

import (
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"strconv"
)

func GetCateList(c *gin.Context) {
	list, err := models.GetAllCate()
	if err != nil {
		panic(err)
	}

	utils.GetCommonResponse(c, list)
}

func CreateOrEditCate(c *gin.Context) {
	cate := models.Cate{}
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		panic(err)
	}

	if cate.Id != 0 {
		ca := models.GetCate(cate.Id)
		if ca == nil {
			utils.GetMessageError("CATE_NOT_FOUND", "Cate not found", c)
			return
		}
	}

	models.CreateOrEditCate(cate)
	utils.GetCommonResponse(c, cate)
}

func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	models.DeleteCate(int32(id))
	utils.GetCommonSuccess(c)
}

func GetCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cate := models.GetCate(int32(id))
	if cate == nil {
		utils.GetMessageError("CATE_NOT_FOUND", "Cate not found", c)
		return
	}
	utils.GetCommonResponse(c, cate)
}
