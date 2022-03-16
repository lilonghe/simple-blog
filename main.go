package main

import (
	"path/filepath"
	"simple-blog/pkg/models"

	"simple-blog/pkg/global"
	"simple-blog/pkg/routers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	if global.Config.Dev {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	// 加载站点配置
	models.LoadOption()

	r := gin.Default()
	theme := global.Options["theme"]
	r.HTMLRender = loadTemplates("./themes/" + theme)
	r.GET("/", routers.Index)
	r.GET("/post/:pathname", routers.PostDetail)
	r.GET("/page/:pathname", routers.PageDetail)
	r.GET("/archives", routers.Archives)
	r.GET("/cate/wiki", routers.Wiki)

	r.Static("/assets", "./themes/"+theme+"/res")

	r.Run()
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
