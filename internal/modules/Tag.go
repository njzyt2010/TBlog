package modules

import "time"

// 标签
type Tag struct {
	Model
	Name string `gorm:"column:name_" json:"name"`
}

func (t Tag) TableName() string {
	return "t_tag"
}

func NewTag() *Tag {
	var curTime = time.Now()
	return &Tag{
		Model:Model{
			CreatedTime: curTime,
			UpdateTime:  curTime,
			Deleted:     false,
		},
	}
}
func UpdateTag() *Tag {
	return &Tag{
		Model: Model{UpdateTime: time.Now()},
	}
}