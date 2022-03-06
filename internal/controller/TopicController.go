package controller

import (
	"TBlog/internal/modules"
	"TBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TopicController struct {
}

func NewTopicController() TopicController {
	return TopicController{}
}

func (t TopicController) Insert(c *gin.Context) {
	json := modules.NewTopic()
	c.BindJSON(&json)

	if err := service.InsertTopic(json); err != nil {
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

func (t TopicController) Update(c *gin.Context) {
	json := modules.NewTopic()
	c.BindJSON(&json)

	if err := service.UpdateTopic(json); err != nil {
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
func (t TopicController) Delete(c *gin.Context) {
	data := make(map[string][]uint)
	c.BindJSON(&data)
	if err := service.DeleteTopicByIds(data["ids"]); err != nil {
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

func (t TopicController) GetById(c *gin.Context) {
	if id ,err:= strconv.ParseUint( c.Query("id"),10,64);err !=nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "400",
			"msg":    "查询错误",
			"result": nil,
		})
	}else {
		topic := service.GetTopicById(uint(id))
		c.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"msg":    "查询成功",
			"result": topic,
		})
	}
}

func (t TopicController) Published(c *gin.Context) {
	json := modules.NewTopic()
	c.BindJSON(&json)

	if err := service.PublishedByTopic(json);err !=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":"400",
			"msg":"发布失败",
			"result":nil,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":"200",
			"msg":"发布成功",
			"result":nil,
		})
	}
}
