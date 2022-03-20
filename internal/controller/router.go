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
	//menu := admin.MenuController
	// 栏目
	topic := admin.TopicController
	// 文章
	article := admin.ArticleController
	// 标签
	tag := admin.TagController
	adminV1 := engine.Group("/v1/admin")
	{
		//adminV1.POST("/menu/insert", menu.Insert)
		//adminV1.POST("/menu/update", menu.Update)
		//adminV1.POST("/menu/delete", menu.DeleteByIds)
		// 栏目
		adminV1.GET("/topic/getById", topic.GetById)
		adminV1.POST("/topic/insert", topic.Insert)
		adminV1.POST("/topic/update", topic.Update)
		adminV1.POST("/topic/delete", topic.Delete)
		//adminV1.POST("/topic/pubAndUnpub", topic.Published)
		// 文章
		adminV1.GET("/article/getById", article.GetById)
		adminV1.POST("/article/insert", article.Insert)
		adminV1.POST("/article/update", article.Update)
		adminV1.POST("/article/delete", article.Delete)
		adminV1.POST("/article/pubSub", article.Published)
		//	标签
		adminV1.POST("/tag/insert", tag.Insert)
		adminV1.POST("/tag/update", tag.Update)
		adminV1.POST("/tag/delete", tag.Delete)
	}
}

//公开访问的部分
func apiEngine(engine *gin.Engine) {
	//专题
	topic := api.TopicController
	apiV1 := engine.Group("/v1/api")
	{
		apiV1.GET("/topics", topic.GetTopicsOfBlog)  // 查询可见专题
		apiV1.GET("/topic/list", topic.GetByTopicId) // 通过专题查询专题下文章
	}
}
