package modules

import (
	"TBlog/pkg/database"
)

// 栏目

type Topic struct {
	Model
	Name      string `gorm:"column:name_" json:"name"`           // 栏目名称
	ImageUrl  string `gorm:"column:image_url" json:"imageUrl"`   //栏目图片
	Type      uint8  `gorm:"column:type_" json:"type"`           // 栏目的类型，0=普通，1=轮播
	Sorted    bool   `gorm:"column:sorted_" json:"sorted"`       // 排序
	Published bool   `gorm:"column:published_" json:"published"` //是否发布
}

func (t Topic) TableName() string {
	return "t_topic"
}
func NewTopic() Topic  {
	return Topic{}
}
// 新增栏目
func (t Topic) Insert() (*Topic, error) {
	if err := database.DB.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

//修改栏目
func (t Topic) Update(values interface{}) (*Topic, error) {
	if err := database.DB.Model(&t).Where("id = ?", t.Id).Updates(values).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

//删除栏目
func (t Topic) DeleteById(ids interface{}) error {
	if err := database.DB.Exec("UPDATE t_topic SET deleted_ = 1 WHERE id in (?)", ids).Error; err != nil {
		return err
	}
	return nil
}

func (t Topic) GetById(id uint) (*Topic,error)  {
	if err := database.DB.First(&t,id).Error;err !=nil {
		return nil,err
	}
	return &t , nil
}

func (t Topic) PublishedById() error {
	if err := database.DB.Model(&t).Update("published_",t.Published).Error ; err !=nil {
		return err
	}
	return nil
}