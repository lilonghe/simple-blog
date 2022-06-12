package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"time"
)

func AddComment(c *gin.Context) {
	addData := models.AddCommentViewModel{}
	data, _ := c.GetRawData()
	err := json.Unmarshal(data, &addData)
	if err != nil {
		utils.GetCommonError(err, c)
		return
	}

	if len(data) > 1024*3 {
		utils.GetCommonError(nil, c)
		return
	}
	comment := models.Comment{
		Nickname:   addData.Nickname,
		Site:       addData.Site,
		Email:      addData.Email,
		Msg:        addData.Msg,
		UserAgent:  c.GetHeader("User-Agent"),
		Ip:         c.GetHeader("X-Real-IP"),
		CreateTime: time.Now(),
	}

	if err = models.AddComment(comment); err != nil {
		utils.GetCommonError(err, c)
		return
	}
	utils.GetCommonSuccess(c)
}

func Comments(c *gin.Context) {
	comments, err := models.GetAllAuditComment()
	if err != nil {
		utils.GetCommonError(err, c)
		return
	}
	utils.GetCommonResponse(c, comments)
}
