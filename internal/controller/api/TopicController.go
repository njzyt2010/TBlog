package api

import (
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type topicController struct {
}

func newTopicController() *topicController {
	return &topicController{}
}

var TopicController = newTopicController()

func (t *topicController) GetTopicsOfBlog(c *gin.Context) {
	data := service.TopicService.GetTopicsOfBlog()
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "查询成功",
		"result": data,
	})
}


