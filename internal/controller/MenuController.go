package controller

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MenuController struct {
}

func NewMenuController() MenuController {
	return MenuController{}
}

func (m MenuController) Insert(c *gin.Context) {
	json := modules.Menu{}
	c.BindJSON(&json)

	if _, err := service.InsertMenu(json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "保存失败",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "保存成功",
		})
	}
}
func (m MenuController) Update(c *gin.Context) {
	json := modules.Menu{}
	c.BindJSON(&json)
	json.UpdateTime = time.Now()
	if err := service.UpdateMenu(json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "保存失败",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "保存成功",
		})
	}
}

func (m MenuController) DeleteByIds(c *gin.Context) {
	json := make(map[string][]uint)
	c.BindJSON(&json)

	if err := service.DeleteMenu(json["id"]); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "保存失败",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "保存成功",
		})
	}
}

func (m MenuController) GetMenusOfBlog(c *gin.Context) {
	data := service.GetMenusOfBlog()
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "查询成功",
		"result": data,
	})
}
