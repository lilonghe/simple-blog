package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"path/filepath"
	"simple-blog/pkg/middlewares"
	"simple-blog/pkg/models"
	"time"

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
	r.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"content-type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	store := cookie.NewStore([]byte(global.Options["password_salt"]))
	r.Use(sessions.Sessions("SESSIONID", store))

	theme := global.Options["theme"]
	r.HTMLRender = loadTemplates("./themes/" + theme)
	r.GET("/", routers.Index)
	r.GET("/post/:pathname", routers.PostDetail)
	r.GET("/page/:pathname", routers.PageDetail)
	r.GET("/archives", routers.Archives)
	r.POST("/comment", routers.AddComment)
	r.GET("/comments", routers.Comments)

	r.Static("/assets", "./themes/"+theme+"/res")

	r.POST("/api/admin/login", routers.Login)
	authorized := r.Group("/api/admin")
	authorized.Use(middlewares.AuthRequired())
	{
		authorized.GET("/user", routers.GetCurrentUser)
		authorized.GET("/system", routers.Dashboard)
	}

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
