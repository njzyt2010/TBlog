package api

import (
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type topicController struct {

}

func newTopicController() *topicController {
	return &topicController{}
}
var TopicController = newTopicController()

func (t *topicController) GetByTopicId(c *gin.Context)  {
	topicId,_ := strconv.Atoi(c.Query("id"))
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	curPage ,_:=strconv.Atoi(c.Query("curPage"))

	if pageSize == 0{
		pageSize = 10
	}
	if curPage == 0 {
		curPage = 1
	}

	articles,total:=service.ArticleService.GetByPage(uint64(topicId),curPage,pageSize)
	result := make(map[string]interface{})
	result["list"] = articles
	result["total"] = total
	c.JSON(http.StatusOK,gin.H{
		"code":"200",
		"msg":"请求成功" ,
		"result":result,
	})
}
