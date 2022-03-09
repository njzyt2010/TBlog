package admin

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagController struct {
}

func NewTagController() TagController {
	return TagController{}
}

func (t TagController) Insert(c *gin.Context) {
	json := modules.NewTag()
	c.BindJSON(&json)
	if err := service.InsertTag(json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "400",
			"msg":  "保存失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "保存成功",
		})
	}
}

func (t TagController) Update(c *gin.Context) {
	json := modules.NewTag()
	c.BindJSON(&json)
	if err := service.UpdateTag(json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "400",
			"msg":  "保存失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "保存成功",
		})
	}
}
func (t TagController) Delete(c *gin.Context) {
	data := make(map[string][]uint)
	c.BindJSON(&data)
	if err := service.DeleteTag(data["ids"]); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "400",
			"msg":  "删除失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "删除成功",
		})
	}
}
