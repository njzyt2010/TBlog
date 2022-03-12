package admin

import (
	"TBlog/internal/modules"
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

func (t *topicController) Insert(c *gin.Context) {
	json := modules.NewTopic()
	c.BindJSON(json)

	if err := service.TopicService.Insert(json); err != nil {
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

func (t *topicController) Update(c *gin.Context) {
	json := modules.UpdateTopic()
	c.BindJSON(json)

	if err := service.TopicService.Updates(json); err != nil {
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
func (t *topicController) Delete(c *gin.Context) {
	data := make(map[string][]uint64)
	c.BindJSON(&data)
	if err := service.TopicService.Delete(data["ids"]); err != nil {
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

func (t *topicController) GetById(c *gin.Context) {
	if id, err := strconv.ParseUint(c.Query("id"), 10, 64); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "查询错误",
			"result": nil,
		})
	} else {
		topic, err2 := service.TopicService.GetById(id)
		if err2 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   "400",
				"msg":    "查询错误",
				"result": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   "200",
				"msg":    "查询成功",
				"result": topic,
			})
		}
	}
}