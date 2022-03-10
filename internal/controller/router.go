package controller

import (
	"TBlog/internal/controller/admin"
	"TBlog/internal/controller/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	adminEngine(engine)
	apiEngine(engine)
	return engine
}

//管理端的 controller
func adminEngine(engine *gin.Engine) {
	//菜单
	menu := admin.NewMenuController()
	// 栏目
	topic := admin.NewTopicController()
	// 文章
	article := admin.NewArticleController()
	// 标签
	tag := admin.NewTagController()
	apiV1 := engine.Group("/v1/admin")
	{
		apiV1.POST("/menu/insert", menu.Insert)
		apiV1.POST("/menu/update", menu.Update)
		apiV1.POST("/menu/delete", menu.DeleteByIds)
		// 栏目
		apiV1.GET("/topic/getById", topic.GetById)
		apiV1.POST("/topic/insert", topic.Insert)
		apiV1.POST("/topic/update", topic.Update)
		apiV1.POST("/topic/delete", topic.Delete)
		apiV1.POST("/topic/pubAndUnpub", topic.Published)
		// 文章
		apiV1.GET("/article/getById", article.GetById)
		apiV1.POST("/article/insert", article.Insert)
		apiV1.POST("/article/update", article.Update)
		apiV1.POST("/article/delete", article.Delete)
		apiV1.POST("/article/pubSub", article.Published)
		//	标签
		apiV1.POST("/tag/insert", tag.Insert)
		apiV1.POST("/tag/update", tag.Update)
		apiV1.POST("/tag/delete", tag.Delete)
	}
}

//公开访问的部分
func apiEngine(engine *gin.Engine) {
	//菜单
	menu := api.MenuController
	apiV1 := engine.Group("/v1/api")
	{
		apiV1.GET("/menu/gets", menu.GetMenusOfBlog)
	}
}
