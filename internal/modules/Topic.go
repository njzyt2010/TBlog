package modules

import (
	"time"
)

// 栏目

type Topic struct {
	Model
	Pid       uint64 `gorm:"column:pid_" json:"pid"`             //子栏目
	Title      string `gorm:"column:title_" json:"title"`           // 栏目名称
	ImageUrl  string `gorm:"column:image_url" json:"imageUrl"`   //栏目图片
	Rotation  bool  `gorm:"column:rotation_" json:"rotation"`    //是否轮播栏目
	Sorted    uint8  `gorm:"column:sorted_" json:"sorted"`       // 排序
	Published bool   `gorm:"column:published_" json:"published"` //是否发布
}

func (t Topic) TableName() string {
	return "t_topic"
}
func NewTopic() *Topic {
	var curTime = time.Now()
	return &Topic{
		Model: Model{
			CreatedTime: curTime,
			UpdateTime:  curTime,
			Deleted:     false,
		},
	}
}

func UpdateTopic() *Topic {
	return &Topic{Model: Model{
		UpdateTime: time.Now(),
		Deleted:    false,
	}}
}