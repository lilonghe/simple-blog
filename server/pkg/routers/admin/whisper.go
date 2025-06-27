package adminRouters

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"simple-blog/pkg/global"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"fmt"
)

// 获取 whisper 列表（分页）
func GetWhisperList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	var isPublic *bool
	if v := c.Query("is_public"); v != "" {
		val, _ := strconv.ParseBool(v)
		isPublic = &val
	}

	whispers, total, err := models.GetWhisperListAdmin(limit, offset, isPublic)
	if err != nil {
		utils.GetCommonError(err, c)
		return
	}
	utils.GetPageResponse(c, whispers, total)
}

// 新建 whisper
func CreateWhisper(c *gin.Context) {
	var whisper models.Whisper
	if err := c.ShouldBindJSON(&whisper); err != nil {
		utils.GetMessageError("PARAM_ERROR", err.Error(), c)
		return
	}
	whisper.CreateTime = time.Now()
	whisper.UpdateTime = time.Now()
	if err := models.CreateWhisper(&whisper); err != nil {
		utils.GetCommonError(err, c)
		return
	}
	utils.GetCommonResponse(c, whisper)
}

// 编辑 whisper
func EditWhisper(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	whisper, err := models.GetWhisperById(int32(id))
	if err != nil || whisper == nil {
		utils.GetMessageError("NOT_FOUND", "not found", c)
		return
	}
	before, _ := json.Marshal(whisper)
	if err := c.ShouldBindJSON(whisper); err != nil {
		utils.GetMessageError("PARAM_ERROR", err.Error(), c)
		return
	}
	fmt.Println(whisper.IsPublic)
	whisper.UpdateTime = time.Now()
	if err := models.UpdateWhisper(int32(id), whisper); err != nil {
		utils.GetCommonError(err, c)
		return
	}
	after, _ := json.Marshal(whisper)
	log := models.OperationLog{
		OpType:        "edit_whisper",
		OpTargetId:    int32(id),
		OpTime:        time.Now(),
		BeforeContent: string(before),
		AfterContent:  string(after),
		OpDetail:      "编辑whisper内容",
	}
	global.Store.Insert(&log)
	utils.GetCommonResponse(c, whisper)
}

// 删除 whisper
func DeleteWhisper(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	whisper, err := models.GetWhisperById(int32(id))
	if err != nil || whisper == nil {
		utils.GetMessageError("NOT_FOUND", "not found", c)
		return
	}
	before, _ := json.Marshal(whisper)
	if err := models.DeleteWhisper(int32(id)); err != nil {
		utils.GetCommonError(err, c)
		return
	}
	log := models.OperationLog{
		OpType:        "delete_whisper",
		OpTargetId:    int32(id),
		OpTime:        time.Now(),
		BeforeContent: string(before),
		AfterContent:  "{}",
		OpDetail:      "删除whisper",
	}
	global.Store.Insert(&log)
	utils.GetCommonSuccess(c)
}

// 修改公开性（不更新更新时间）
func UpdateWhisperVisibility(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	whisper, err := models.GetWhisperById(int32(id))
	if err != nil || whisper == nil {
		fmt.Println(err)
		utils.GetMessageError("NOT_FOUND", "not found", c)
		return
	}
	before, _ := json.Marshal(whisper)
	type req struct {
		IsPublic bool `json:"is_public"`
	}
	var body req
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.GetMessageError("PARAM_ERROR", err.Error(), c)
		return
	}
	if err := models.UpdateWhisperVisibility(int32(id), body.IsPublic); err != nil {
		utils.GetCommonError(err, c)
		return
	}
	whisper.IsPublic = body.IsPublic
	after, _ := json.Marshal(whisper)
	log := models.OperationLog{
		OpType:        "update_whisper_visibility",
		OpTargetId:    int32(id),
		OpTime:        time.Now(),
		BeforeContent: string(before),
		AfterContent:  string(after),
		OpDetail:      "修改whisper公开性",
	}
	global.Store.Insert(&log)
	utils.GetCommonResponse(c, whisper)
} 