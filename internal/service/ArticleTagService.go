package service

import (
	"TBlog/internal/modules"
	"TBlog/internal/repository"
)

type articleTagService struct {
}

func newArticleTagService() *articleService {
	return &articleService{}
}

var ArticleTagService = newArticleService()

func (t *articleTagService) Insert(at *modules.ArticleTag) error {
	return repository.ArticleTagRepository.Insert(at)
}
func (t *articleTagService) DeleteByArticleId(aid uint64) error {
	return repository.ArticleTagRepository.DeleteByArticleId(aid)
}
