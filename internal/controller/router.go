package controller

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//菜单
	menu := NewMenuController()
	// 栏目
	topic := NewTopicController()
	// 文章
	article := NewArticleController()
	// 标签
	tag := NewTagController()
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/menu/insert", menu.Insert)
		apiV1.POST("/menu/update", menu.Update)
		apiV1.POST("/menu/delete", menu.DeleteByIds)
		apiV1.GET("/menu/gets", menu.GetMenusOfBlog)
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
		apiV1.POST("/article/pubAndUnpub", article.Published)
		//	标签
		apiV1.POST("/tag/insert", tag.Insert)
		apiV1.POST("/tag/update", tag.Update)
		apiV1.POST("/tag/delete", tag.Delete)
	}

	return router
}
