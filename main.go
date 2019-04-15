package main

import (
	"github.com/lilonghe/simple-blog/pkg/models"
	"path/filepath"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/lilonghe/simple-blog/pkg/global"
	"github.com/lilonghe/simple-blog/pkg/routers"
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
	r.GET("/archives", routers.Archives)

	r.Static("/assets", "./themes/"+theme+"/res")

	go tickerRefreshOption()
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

// 因为暂无控制面板，所以暂时为定时刷新主题配置
func tickerRefreshOption() {
	d := time.Duration(time.Minute * 5)

	t := time.NewTicker(d)
	defer t.Stop()

	for {
		<-t.C
		models.LoadOption()
	}
}
