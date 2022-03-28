package repository

import (
	"TBlog/internal/modules"
	"TBlog/pkg/database"
	"time"
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
	err := database.DB.Exec("update t_article set published_ = ? ,published_time = ? where id=?",value,time.Now(),id).Error
	return err
}
//通过栏目查询文章
func (a *articleRepository) GetByPage(topicId uint64,curPage int,pageSize int) ([]modules.Article, int64)  {
	var total int64 = 0
	database.DB.Model(&modules.Article{}).Where( "topic_id = ? and published_ =? and deleted_ = ?", topicId,true,false).Count(&total)
	var artiles []modules.Article = nil

	database.DB.Model(&modules.Article{}).Where("topic_id = ? and published_ =? and deleted_ = ?", topicId,true,false).Offset((curPage - 1)* pageSize).Limit(pageSize).Find(&artiles)
	return artiles,total
}
// 通过 tag查询文章
func (a *articleRepository) GetByTagIdPage(tagId uint64,curPage int,pageSize int)  ([]modules.Article, int64) {
	var sqlAppend = " FROM t_article ta WHERE ta.deleted_ =0 AND ta.published_ =1 AND ta.id IN ( SELECT tat.article_id  FROM t_article_tag tat "

	if tagId > 0 {
		sqlAppend +=" WHERE tat.tag_id =?"
	}
	sqlAppend += ") ORDER BY ta.published_time DESC"

	var total int64 =0
	var artiles []modules.Article = nil

	if tagId > 0 {
		database.DB.Raw("SELECT count(DISTINCT ta.id) " + sqlAppend,tagId).Scan(&total)
		database.DB.Raw("select DISTINCT * "+sqlAppend+" limit ?,?", tagId,(curPage - 1)* pageSize,pageSize).Scan(&artiles)
	}else {
		database.DB.Raw("SELECT count(DISTINCT ta.id) " + sqlAppend).Scan(&total)
		database.DB.Raw("select DISTINCT ta.id,ta.title_,ta.update_time,ta.topic_id "+sqlAppend+" limit ?,?",(curPage - 1)* pageSize,pageSize).Scan(&artiles)
	}

	return artiles,total
}