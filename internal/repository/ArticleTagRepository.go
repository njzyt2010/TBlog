package repository

import (
	"TBlog/internal/modules"
	"TBlog/pkg/database"
)

type articleTagRepository struct {
}

func newArticleTagRepository() *articleTagRepository {
	return &articleTagRepository{}
}

var ArticleTagRepository = newArticleTagRepository()

func (a *articleTagRepository) Insert(at *modules.ArticleTag) error {
	err := database.DB.Create(at).Error
	return err
}
func (a *articleTagRepository) DeleteByArticleId(aid uint64) error {
	err := database.DB.Model(&modules.ArticleTag{}).Where("article_id = ?", aid).Error
	return err
}

func (a *articleTagRepository) GetByPage(tagId uint64,curPage int ,pageSize int) ([]modules.ArticleTag,int64)  {
	var total int64 = 0
	database.DB.Model(&modules.ArticleTag{}).Where( "tag_id = ?",tagId).Count(&total)

	var articleTags []modules.ArticleTag = nil
	database.DB.Model(&modules.ArticleTag{}).Where("tag_id = ?",tagId).Find(&articleTags).Offset((curPage - 1)* pageSize).Limit(pageSize).Find(&articleTags)

	return articleTags,total
}