package router

import (
	"TBlog/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//菜单
	menu := controller.NewMenuController()
	// 栏目
	topic := controller.NewTopicController()
	// 文章
	// 标签

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
		apiV1.POST("/topic/pubAndUnpub" , topic.Published)

	}

	return router
}
