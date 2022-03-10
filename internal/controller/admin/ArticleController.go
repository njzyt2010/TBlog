package admin

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type articleController struct {
}

func newArticleController() *articleController {
	return &articleController{}
}

var ArticleController = newArticleController()

func (a *articleController) Insert(c *gin.Context) {
	json := modules.NewArticle()
	c.BindJSON(json)

	if err := service.ArticleService.Insert(json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "保存失败",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"msg":    "保存成功",
			"result": nil,
		})
	}
}

func (a *articleController) Update(c *gin.Context) {
	json := modules.UpdateArticle()
	c.BindJSON(json)

	if err := service.ArticleService.Update(json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "保存失败",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"msg":    "保存成功",
			"result": nil,
		})
	}
}
func (a *articleController) Delete(c *gin.Context) {
	json := make(map[string][]uint64)
	c.BindJSON(&json)
	var ids []uint64 = json["ids"]
	if err := service.ArticleService.Delete(ids); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "删除失败",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"msg":    "删除成功",
			"result": nil,
		})
	}
}

func (a *articleController) GetById(c *gin.Context) {
	if id, err := strconv.ParseUint(c.Query("id"), 10, 64); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "查询错误",
			"result": nil,
		})
	} else {
		article := service.ArticleService.GetById(id)
		c.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"msg":    "查询成功",
			"result": article,
		})
	}
}

func (t *articleController) Published(c *gin.Context) {
	json := modules.NewArticle()
	c.BindJSON(json)

	if err := service.ArticleService.PubSubById(json.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "发布失败",
			"result": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"msg":    "发布成功",
			"result": nil,
		})
	}
}
