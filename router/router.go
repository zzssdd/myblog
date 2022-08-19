package router

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"myblog/api"
	"myblog/middleware"
	"myblog/utils"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "./web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	router := r.Group("api")
	{
		router.GET("article/list", api.GetArtList)
		router.GET("article/:id", api.GetArtInfo)
		router.GET("articlecate/:name", api.GetCateArticle)
		router.GET("comment/:id", api.GetCommentList)
		router.POST("comment/add", api.AddComment)
		router.GET("category", api.GetCategory)

		router.GET("life", api.GetLife)
		router.GET("life/:id", api.GetLifeInfo)

		router.GET("lifecomment/:id", api.GetLifecomment)
		router.POST("lifecomment/add", api.AddLcomment)

		router.POST("board/add", api.AddBoard)
		router.GET("board", api.GetBoard)

		router.GET("project", api.GetProject)

		router.POST("login", api.Login)
	}

	auth := r.Group("api")
	auth.Use(middleware.JwtToken())
	{
		auth.PUT("article/edit/:id", api.EditArticle)
		auth.POST("article/add", api.AddArticle)
		auth.GET("article/search", api.SearchArticle)
		auth.DELETE("article/delete/:id", api.DeleteArticle)

		auth.GET("artcomment/:id", api.GetArtComment)
		auth.DELETE("comment/delete/:id", api.DeleteComment)
		auth.GET("artcomment/list", api.GetArtCommentList)

		auth.POST("category/add", api.AddCategory)
		auth.DELETE("category/delete/:id", api.DelCategory)
		auth.PUT("category/:id", api.EditCategory)
		auth.GET("categoryid/:name", api.GetCategoryId)

		auth.POST("life/add", api.AddLife)
		auth.DELETE("life/delete/:id", api.DeleteLife)
		auth.PUT("life/edit/:id", api.EditLife)

		auth.GET("lifecomment/list", api.GetLcomment)
		auth.DELETE("lifecomment/delete/:id", api.DeleteLcomment)

		auth.DELETE("board/delete/:id", api.DeleteBoard)

		auth.GET("project/:id", api.GetProjectInfo)
		auth.POST("project/add", api.AddProject)
		auth.DELETE("project/delete/:id", api.DeleteProject)
		auth.PUT("project/edit/:id", api.EditProject)
	}
	r.Run(utils.Config.HttpPort)
}
