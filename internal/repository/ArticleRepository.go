package repository

import (
	"TBlog/internal/modules"
	"TBlog/pkg/database"
)

type articleRepository struct {
}

var ArticleRepository = newArticleRepository()

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}

// 新增文章
func (a *articleRepository) Insert(article *modules.Article) error {
	err := database.DB.Create(article).Error
	return err
}

// 修改文章
func (a *articleRepository) Update(article *modules.Article) error {
	err := database.DB.Updates(article).Error
	return err
}

// 修改文章
func (a *articleRepository) Updates(id uint64, values map[string]interface{}) error {
	err := database.DB.Model(&modules.Article{}).Where("id = ?", id).Updates(values).Error
	return err
}

// 修改文章
func (a *articleRepository) UpdateColumn(id uint64, column string, value interface{}) error {
	err := database.DB.Model(&modules.Article{}).Where("id = ?", id).UpdateColumn(column, value).Error
	return err
}

// 删除文章
func (a *articleRepository) Delete(id uint64) error {
	if err := database.DB.Model(&modules.Article{}).Where("id = ?" ,id).Update("deleted_" ,1).Error; err != nil {
		return err
	}
	return nil
}

//通过主键查询
func (a *articleRepository) GetById(id uint64) (*modules.Article, error) {
	var article *modules.Article = &modules.Article{}
	if err := database.DB.First(article, id).Error; err != nil {
		return nil, err
	}
	return article, nil
}
func (a *articleRepository) PubSubById(id uint64, value interface{}) error {
	err := a.UpdateColumn(id, "published_", value)
	return err
}

func (a *articleRepository) GetByPage(topicId uint64,curPage int,pageSize int) ([]modules.Article, int64)  {
	var total int64 = 0
	database.DB.Model(&modules.Article{}).Where( "topic_id = ? and published_ =? and deleted_ = ?", topicId,true,false).Count(&total)
	var artiles []modules.Article = nil

	database.DB.Model(&modules.Article{}).Where("topic_id = ? and published_ =? and deleted_ = ?", topicId,true,false).Offset((curPage - 1)* pageSize).Limit(pageSize).Find(&artiles)
	return artiles,total
}
