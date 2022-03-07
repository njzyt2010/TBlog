package modules

import "TBlog/pkg/database"

// 文章
type Article struct {
	Model
	Title     string `gorm:"column:title_" json:"title"`         //文章标题
	Content   string `gorm:"column:content_" json:"content"`     //内容
	Original  bool   `gorm:"column:original_" json:"original"`   // 是否原创文章
	Published bool   `gorm:"column:published_" json:"published"` //文章是否发布
}

func (a Article) TableName() string {
	return "t_article"
}
func  NewArticle() Article {
	return Article{}
}
// 新增文章
func (a Article) Insert() (*Article, error) {
	if err := database.DB.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

// 修改文章
func (a Article) Update(values interface{}) (*Article, error) {
	if err := database.DB.Model(&a).Where("id = ?", a.Id).Updates(values).Error; err != nil {
		return nil, err
	}
	return &a, nil
}
func (a Article) Delete(ids interface{}) error {
	if err := database.DB.Exec("UPDATE t_article SET deleted_=1 where id in (?)", ids).Error; err != nil {
		return err
	}
	return nil
}
func (a Article) GetById(id interface{}) (*Article, error) {
	if err := database.DB.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) PubOrUnpubById() error {
	if err := database.DB.Model(&a).Update("published_",a.Published).Error ; err !=nil {
		return err
	}
	return nil
}
