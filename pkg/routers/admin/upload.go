package adminRouters

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"simple-blog/pkg/global"
	"simple-blog/pkg/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}

	accessUrl := ""
	relativePath := time.Now().Format("20060102")
	fileSuffix := path.Ext(file.Filename)
	saveName := "upload_" + strings.ReplaceAll(uuid.NewString(), "-", "") + fileSuffix

	// get upload config, must config upload platform and CDN host
	if global.Options["upload"] != "" && global.Config.CDNHost != "" {
		ossConfig := make(map[string]string, 0)
		err = json.Unmarshal([]byte(global.Options["upload"]), &ossConfig)
		if err != nil {
			panic(err)
		}

		if ossConfig["type"] == "tencent" {
			accessUrl, err = uploadTencent(ossConfig, file, saveName, relativePath)
		}
	} else {
		accessUrl, err = uploadLocal(c, file, saveName, relativePath)
	}

	if err != nil {
		panic(err)
	}

	utils.GetCommonResponse(c, accessUrl)
}

func uploadTencent(ossConfig map[string]string, file *multipart.FileHeader, saveName string, relativePath string) (string, error) {
	fileReader, err := file.Open()
	if err != nil {
		return "", err
	}

	url, _ := url.Parse(ossConfig["origin"])
	bucket := &cos.BaseURL{BucketURL: url}
	client := cos.NewClient(bucket, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  ossConfig["secretId"],
			SecretKey: ossConfig["secretKey"],
		},
	})

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentLength: file.Size,
		},
	}

	relativeFullPath := ossConfig["prefix"] + "/" + relativePath + "/" + saveName

	_, err = client.Object.Put(context.Background(), relativeFullPath, fileReader, opt)
	if err != nil {
		return "", err
	}

	return global.Config.CDNHost + "/" + relativeFullPath, nil
}

func uploadLocal(c *gin.Context, file *multipart.FileHeader, saveName string, relativePath string) (string, error) {
	utils.AutoCreateFolder(global.Config.UploadPath, relativePath)
	relativeFullPath := global.Config.UploadPath + "/" + relativePath + "/" + saveName
	err := c.SaveUploadedFile(file, path.Join(global.Config.UploadPath, relativeFullPath))
	if err != nil {
		return "", err
	}
	return global.Options["site_url"] + path.Join(global.Config.UploadAccessPath, relativeFullPath), err
}
