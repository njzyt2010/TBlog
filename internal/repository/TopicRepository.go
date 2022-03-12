package repository

import (
	"TBlog/internal/modules"
	"TBlog/pkg/database"
)

type topicRepository struct {
}

func newTopicRepository() *topicRepository {
	return &topicRepository{}
}

var TopicRepository = newTopicRepository()

func (t *topicRepository) Insert(topic *modules.Topic) error {
	err := database.DB.Create(topic).Error
	return err
}

func (t *topicRepository) Updates(topic *modules.Topic) error {
	err := database.DB.Updates(topic).Error
	return err
}

func (t *topicRepository) Update(id uint64, values map[string]interface{}) error {
	err := database.DB.Model(&modules.Topic{}).Where("id = ?", id).Updates(values).Error
	return err
}

func (t *topicRepository) UpdateColumn(id uint64, column string, value interface{}) error {
	err := database.DB.Model(&modules.Topic{}).Where("id = ?", id).UpdateColumn(column, value).Error
	return err
}

func (t *topicRepository) Delete(ids []uint64) error {
	err := database.DB.Model(&modules.Topic{}).Where("id in (?)", ids).Error
	return err
}

func (t *topicRepository) GetById(id uint64) (*modules.Topic, error) {
	var topic *modules.Topic = nil
	if err := database.DB.First(topic, id).Error; err != nil {
		return topic, err
	}
	return topic, nil

}
