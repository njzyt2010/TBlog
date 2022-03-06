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
	// 文章
	// 标签
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/menu/insert", menu.Insert)
		apiV1.POST("/menu/update",menu.Update)
		apiV1.POST("/menu/delete",menu.DeleteByIds)
	}

	return router
}
