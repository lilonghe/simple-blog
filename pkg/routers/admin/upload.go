package adminRouters

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path"
	"simple-blog/pkg/global"
	"simple-blog/pkg/utils"
	"strings"
	"time"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}

	dayDir := time.Now().Format("20060102")
	fileSuffix := path.Ext(file.Filename)
	relativeFullPath := dayDir + "/upload_" + strings.ReplaceAll(uuid.NewString(), "-", "") + fileSuffix
	utils.AutoCreateFolder(global.Config.UploadPath, dayDir)
	err = c.SaveUploadedFile(file, path.Join(global.Config.UploadPath, relativeFullPath))
	if err != nil {
		panic(err)
	}
	utils.GetCommonResponse(c, global.Options["site_url"]+path.Join(global.Config.UploadAccessPath, relativeFullPath))
}
