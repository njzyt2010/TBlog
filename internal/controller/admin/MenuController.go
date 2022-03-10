package admin

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type menuController struct {
}

func newMenuController() *menuController {
	return &menuController{}
}

var MenuController = newMenuController()

func (m *menuController ) Insert(c *gin.Context) {
	json := modules.NewMenu()
	c.BindJSON(json)

	if err := service.MenuService.Insert(json); err != nil {
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
func (m *menuController ) Update(c *gin.Context) {
	json := modules.UpdateMenu()
	c.BindJSON(json)
	json.UpdateTime = time.Now()
	if err := service.MenuService.Update(json); err != nil {
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

func (m *menuController ) DeleteByIds(c *gin.Context) {
	json := make(map[string][]uint64)
	c.BindJSON(&json)
	var id uint64 = json["id"][0]
	if err := service.MenuService.Delete(id); err != nil {
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

