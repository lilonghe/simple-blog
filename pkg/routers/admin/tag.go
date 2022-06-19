package adminRouters

import (
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"strconv"
)

func GetTagList(c *gin.Context) {
	list, err := models.GetAllTagWithCount()
	if err != nil {
		panic(err)
	}

	utils.GetCommonResponse(c, list)
}

func CreateOrEditTag(c *gin.Context) {
	tag := models.Tag{}
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		panic(err)
	}

	if tag.Id != 0 {
		t := models.GetTag(tag.Id)
		if t == nil {
			utils.GetMessageError("TAG_NOT_FOUND", "Tag not found", c)
			return
		}
	}

	repeat := models.CheckCateNameOrPathRepeat(tag.Name, tag.Pathname, tag.Id)
	if repeat {
		utils.GetMessageError("REPEAT_TAG_INFO", "Tag name or pathname repeat", c)
		return
	}

	models.CreateOrEditTag(tag)
	utils.GetCommonResponse(c, tag)
}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	models.DeleteTag(int32(id))
	utils.GetCommonSuccess(c)
}

func GetTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cate := models.GetTag(int32(id))
	if cate == nil {
		utils.GetMessageError("TAG_NOT_FOUND", "Tag not found", c)
		return
	}
	utils.GetCommonResponse(c, cate)
}
