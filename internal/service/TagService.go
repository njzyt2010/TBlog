package service

import (
	"TBlog/internal/modules"
)

func InsertTag(tag modules.Tag) error {
	if _, err := tag.Insert(); err != nil {
		return err
	}
	return nil
}

func UpdateTag(tag modules.Tag) error {

	if _, err := tag.Update(tag); err != nil {
		return err
	}
	return nil
}

func DeleteTag(ids interface{}) error {
	if err := modules.NewTag().Delete(ids); err != nil {
		return err
	}
	return nil
}
