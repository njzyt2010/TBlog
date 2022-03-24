package api

import (
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

func (a *articleController) GetByTopicId(c *gin.Context) {
	topicId, _ := strconv.Atoi(c.Query("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	curPage, _ := strconv.Atoi(c.Query("curPage"))

	if pageSize == 0 {
		pageSize = 10
	}
	if curPage == 0 {
		curPage = 1
	}

	articles, total := service.ArticleService.GetByPage(uint64(topicId), curPage, pageSize)
	result := make(map[string]interface{})
	result["list"] = articles
	result["total"] = total
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "请求成功",
		"result": result,
	})
}
// 通过id查询文章
func (a articleController) GetById(c *gin.Context)  {
	articleId, _ := strconv.Atoi(c.Query("id"))
	article := service.ArticleService.GetById(uint64(articleId)) ;

	c.JSON(http.StatusOK,gin.H{
		"code":"200",
		"msg":"查询成功" ,
		"result":article,
	})
}

func (a *articleController) GetByTag(c *gin.Context)  {
	tagId,_ := strconv.Atoi(c.Query("tagId"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	curPage, _ := strconv.Atoi(c.Query("curPage"))

	if pageSize == 0 {
		pageSize = 10
	}
	if curPage == 0 {
		curPage = 1
	}

	articles,total := service.ArticleService.GetByTagIdPage(uint64(tagId),curPage,pageSize);
	result := make(map[string]interface{})
	result["list"] = articles
	result["total"] = total
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "请求成功",
		"result": result,
	})
}

//查询最新数据
func (a *articleController) GetNewer(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	curPage, _ := strconv.Atoi(c.Query("curPage"))

	if pageSize == 0 {
		pageSize = 10
	}
	if curPage == 0 {
		curPage = 1
	}

	articles, total := service.ArticleService.GetNewer( curPage, pageSize)
	result := make(map[string]interface{})
	result["list"] = articles
	result["total"] = total
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "请求成功",
		"result": result,
	})
}
