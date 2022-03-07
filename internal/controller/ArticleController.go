package controller

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
	c.BindJSON(&json)

	if err := service.InsertArticle(json); err != nil {
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
	c.BindJSON(&json)

	if err := service.UpdateArticle(json); err != nil {
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
	data := make(map[string][]uint)
	c.BindJSON(&data)
	if err := service.DeleteArticleById(data["ids"]); err != nil {
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
		article, _ := service.GetArticleById(uint(id))
		c.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"msg":    "查询成功",
			"result": article,
		})
	}
}

func (t ArticleController) Published(c *gin.Context) {
	json := modules.NewArticle()
	c.BindJSON(&json)

	if err := service.PubOrUnpubArticle(json); err != nil {
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
