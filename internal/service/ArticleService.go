package service

import (
	"TBlog/internal/modules"
	"TBlog/internal/repository"
)

type articleService struct {
}

func newArticleService() *articleService {
	return &articleService{}
}

var ArticleService = newArticleService()

//新增
func (a *articleService) Insert(article *modules.Article) error {
	return repository.ArticleRepository.Insert(article)
}

// 修改
func (a *articleService) Update(article *modules.Article) error {
	return repository.ArticleRepository.Update(article)
}
func (a *articleService) Updates(id uint64, values map[string]interface{}) error {
	return repository.ArticleRepository.Updates(id, values)
}
func (a *articleService) UpdateColumn(id uint64, column string, value interface{}) error {
	return repository.ArticleRepository.UpdateColumn(id, column, value)
}

//删除
func (a *articleService) Delete(ids []uint64) error {

	for i := 0; i < len(ids); i++ {
		if err := repository.ArticleRepository.Delete(ids[i]); err != nil {
			return err
		}
	}

	return nil
}
func (a *articleService) GetById(id uint64) *modules.Article {
	if data, err := repository.ArticleRepository.GetById(id); err != nil {
		return nil
	} else {
		return data
	}
}

//发布、取消发布
func (a *articleService) PubSubById(id uint64) error {
	if article, err := repository.ArticleRepository.GetById(id); err != nil {
		return err
	} else {
		return repository.ArticleRepository.PubSubById(id, !*article.Published)
	}
}

func (a *articleService) GetByPage(topicId uint64,curPage int ,pageSize int) ([]modules.Article,int64)  {
	return repository.ArticleRepository.GetByPage(topicId,curPage,pageSize)
}

func (a *articleService) GetByTagIdPage(tagId uint64,curPage int ,pageSize int) ([]modules.Article,int64) {
	return repository.ArticleRepository.GetByTagIdPage(tagId,curPage,pageSize)
}
// 查询最近更新的文章
func (a *articleService) GetNewer(curPage int,pageSize int) ([]modules. Article,int64){
	return repository.ArticleRepository.GetNewer(curPage,pageSize)
}