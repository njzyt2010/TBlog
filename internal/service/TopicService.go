package service

import (
	"TBlog/internal/modules"
	"time"
)

func InsertTopic(t modules.Topic) error {
	var curTime = time.Now()
	t.CreatedTime = curTime
	t.UpdateTime = curTime
	if _, err := t.Insert(); err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateTopic(t modules.Topic) error {
	t.UpdateTime = time.Now()
	if _, err := t.Update(t); err != nil {
		return err
	} else {
		return nil
	}
}
func DeleteTopicByIds(ids []uint) error {
	if err := modules.NewTopic().DeleteById(ids); err != nil {
		return err
	}
	return nil
}

func GetTopicById(id uint) *modules.Topic {
	if data, err := modules.NewTopic().GetById(id); err != nil {
		return nil
	} else {
		return data
	}
}

func PublishedByTopic(t modules.Topic) error {
	if err := t.PublishedById();err !=nil {
		return err
	}
	return nil
}
