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
