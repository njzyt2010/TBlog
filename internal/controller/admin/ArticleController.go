package admin

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleController struct {
}

func NewArticleController() ArticleController {
	return ArticleController{}
}

func (a ArticleController) Insert(c *gin.Context) {
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

func (a ArticleController) Update(c *gin.Context) {
	json := modules.NewArticle()
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
func (a ArticleController) Delete(c *gin.Context) {
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

func (a ArticleController) GetById(c *gin.Context) {
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

func (t ArticleController) Published(c *gin.Context) {
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
