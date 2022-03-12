package admin

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tagController struct {
}

func newTagController() *tagController {
	return &tagController{}
}

var TagController = newTagController()

func (t *tagController) Insert(c *gin.Context) {
	json := modules.NewTag()
	c.BindJSON(json)
	if err := service.TagService.Insert(json); err != nil {
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

func (t *tagController) Update(c *gin.Context) {
	json := modules.UpdateTag()
	c.BindJSON(json)
	if err := service.TagService.Updates(json); err != nil {
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
func (t *tagController) Delete(c *gin.Context) {
	json := make(map[string][]uint64)
	c.BindJSON(json)
	if err := service.TagService.Deletes(json["ids"]); err != nil {
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

//
func (t *tagController) GetAll(c *gin.Context) {
	data := service.TagService.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "删除成功",
		"result": data,
	})
}
