package service

import (
	"TBlog/internal/modules"
	"TBlog/internal/repository"
)

type tagService struct {
}

func newTagService() *tagService {
	return &tagService{}
}

var TagService = newTagService()

func (t *tagService) Insert(tag *modules.Tag) error {
	return repository.TagRepository.Insert(tag)
}

func (t *tagService) Updates(tag *modules.Tag) error {
	return repository.TagRepository.Updates(tag)
}
func (t *tagService) Deletes(ids []uint64) error {
	return repository.TagRepository.Deletes(ids)
}
func (t *tagService) DeleteByName(name string) error {
	return repository.TagRepository.DeleteByName(name)
}

func (t *tagService) GetByIds(ids []uint64) []modules.Tag {
	return repository.TagRepository.GetByIds(ids)
}

func (t *tagService) GetAll() []modules.Tag {
	return repository.TagRepository.GetAll()
}
