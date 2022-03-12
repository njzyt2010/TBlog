package repository

import (
	"TBlog/internal/modules"
	"TBlog/pkg/database"
)

type tagRepository struct {
}

func newTagRepository() *tagRepository {
	return &tagRepository{}
}

var TagRepository = newTagRepository()

func (t *tagRepository) Insert(tag *modules.Tag) error {
	err := database.DB.Create(tag).Error
	return err
}

func (t *tagRepository) Updates(tag *modules.Tag) error {
	err := database.DB.Updates(tag).Error
	return err
}
func (t *tagRepository) Deletes(ids []uint64) error  {
	err := database.DB.Model(&modules.Tag{}).Where("id in(?)",ids).UpdateColumn("deleted_",true).Error
	return err
}
func (t *tagRepository) DeleteByName(name string) error {
	err := database.DB.Model(&modules.Tag{}).Where("name_ = ?", name).Update("deleted_", true).Error
	return err
}

func (t *tagRepository) GetByIds(ids []uint64) []modules.Tag {
	var tags []modules.Tag = nil
	if err := database.DB.Model(&modules.Tag{}).Where("id in (?)", ids).Find(&tags).Error; err != nil {
		return nil
	}
	return tags
}

func (t *tagRepository) GetAll() []modules.Tag {
	var tags []modules.Tag = nil
	if err := database.DB.Model(&modules.Tag{}).Where("deleted_", false).Find(&tags).Error; err != nil {
		return nil
	}
	return tags
}
