package main

import (
	"net/http"
	"path/filepath"
	"simple-blog/pkg/middlewares"
	"simple-blog/pkg/models"
	adminRouters "simple-blog/pkg/routers/admin"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

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
		AllowMethods:     []string{"DELETE"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	store := cookie.NewStore([]byte(global.Options["password_salt"]))
	store.Options(sessions.Options{MaxAge: 86400, HttpOnly: true, Secure: true, SameSite: http.SameSiteLaxMode})
	r.Use(sessions.Sessions("SESSIONID", store))

	theme := global.Options["theme"]
	r.HTMLRender = loadTemplates("./themes/" + theme)
	r.GET("/", routers.Index)
	r.GET("/post/:pathname", routers.PostDetail)
	r.GET("/page/:pathname", routers.PageDetail)
	r.GET("/archives", routers.Archives)
	r.GET("/sitemap.txt", routers.SitemapTxt)

	r.Static("/assets", "./themes/"+theme+"/res")
	r.Static(global.Config.UploadAccessPath, global.Config.UploadPath)

	if global.Config.EnableAdmin {
		r.POST("/api/admin/login", adminRouters.Login)
		authorized := r.Group("/api/admin")
		authorized.Use(middlewares.AuthRequired())
		{
			authorized.GET("/user", adminRouters.GetCurrentUser)
			authorized.GET("/system", adminRouters.Dashboard)
			authorized.GET("/options", adminRouters.GetOptions)
			authorized.POST("/options", adminRouters.UpdateOptions)
			authorized.POST("/logout", adminRouters.Logout)
			authorized.POST("/upload", adminRouters.Upload)
			authorized.GET("/visit/list", adminRouters.GetVisitList)

			authorized.GET("/post/list", adminRouters.GetPostList)
			authorized.POST("/post", adminRouters.CreateOrEditPost)
			authorized.GET("/post/:id", adminRouters.GetEditPost)
			authorized.DELETE("/post/:id", adminRouters.DeletePost)

			authorized.GET("/cate/list", adminRouters.GetCateList)
			authorized.POST("/cate", adminRouters.CreateOrEditCate)
			authorized.GET("/cate/:id", adminRouters.GetCate)
			authorized.DELETE("/cate/:id", adminRouters.DeleteCate)

			authorized.GET("/tag/list", adminRouters.GetTagList)
			authorized.POST("/tag", adminRouters.CreateOrEditTag)
			authorized.GET("/tag/:id", adminRouters.GetTag)
			authorized.DELETE("/tag/:id", adminRouters.DeleteTag)
		}
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

	if len(layouts) == 0 || len(includes) == 0 {
		panic("can not find theme: " + templatesDir)
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
