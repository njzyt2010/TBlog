package modules

import (
	"TBlog/pkg/database"
	"time"
)

// 标签
type Tag struct {
	Model
	Name string `gorm:"column:name_" json:"name"`
}

func (t Tag) TableName() string {
	return "t_tag"
}

func  NewTag() Tag {
	return Tag{}
}

//新增 标签
func (t Tag) Insert() (*Tag, error) {
	var curTime = time.Now()
	t.CreatedTime = curTime
	t.UpdateTime = curTime
	if err := database.DB.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

// 修改标签
func (t Tag) Update(values interface{}) (*Tag, error) {
	t.UpdateTime = time.Now()
	if err := database.DB.Model(&t).Updates(values).Where("id = ?", t.Id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}
// 删除标签
func (t Tag) Delete(ids interface{}) error {
	if err := database.DB.Exec("UPDATE t_tag SET deleted_ = 1 WHERE id IN (?)", ids).Error; err != nil {
		return err
	}
	return nil
}
