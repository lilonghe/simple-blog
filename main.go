package main

import (
	"path/filepath"

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

	r := gin.Default()
	theme := "default"
	r.HTMLRender = loadTemplates("./themes/" + theme)
	r.GET("/", routers.Index)
	r.Static("/static", "./themes/"+theme+"/res")

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
