package api

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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


// 通过tagId查询最新文章，如果tagId为空，则为查询最新发布
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

func (a *articleController) GetLastAndNext(c *gin.Context) {
	id,_ := strconv.Atoi( c.Query("id"))
	if id > 0 {
		var ret = make(map[string]modules.Article,2)
		last,next := service.ArticleService.GetLastAndNextArticle(uint64(id))
		ret["last"] = last 
		ret["next"] = next
		c.JSON(http.StatusOK,gin.H{
			"code":"200" ,
			"msg":"请求成功",
			"result": ret,
		})
	}
}