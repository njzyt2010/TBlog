package service

import (
	"TBlog/internal/modules"
	"TBlog/internal/repository"
)

type topicService struct {

}

func newTopicService() *topicService {
	return &topicService{}
}
var TopicService = newTopicService()

func (t *topicService) Insert(topic *modules.Topic) error {
	return repository.TopicRepository.Insert(topic)
}

func (t *topicService) Updates(topic *modules.Topic) error {
	return repository.TopicRepository.Updates(topic)
}

func (t *topicService) Update(id uint64,values map[string]interface{}) error  {
	return repository.TopicRepository.Update(id,values)
}
func (t *topicService) UpdateColumn(id uint64,column string,value interface{})error  {
	return repository.TopicRepository.UpdateColumn(id ,column,value)
}
func (t *topicService) Delete(ids []uint64) error {
	return repository.TopicRepository.Delete(ids)
}
func (t *topicService) GetById(id uint64) (*modules.Topic,error) {
	return repository.TopicRepository.GetById(id)
}