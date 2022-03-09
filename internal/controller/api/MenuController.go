package api

import (
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type menuController struct {
}

func newMenuController() *menuController {
	return &menuController{}
}

var MenuController = newMenuController()

// 查询博客的全部菜单
func (m *menuController) GetMenusOfBlog(c *gin.Context) {
	data := service.MenuService.GetMenusOfBlog()
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "查询成功",
		"result": data,
	})
}
