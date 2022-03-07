package service

import (
	"TBlog/internal/modules"
	"time"
)

func InsertArticle(a modules.Article) error {
	var curTime = time.Now()
	a.CreatedTime = curTime
	a.UpdateTime = curTime
	if _, err := a.Insert(); err != nil {
		return err
	}
	return nil
}

func UpdateArticle(a modules.Article) error {
	a.UpdateTime = time.Now()
	if _, err := a.Update(a); err != nil {
		return err
	}
	return nil
}

func DeleteArticleById(ids interface{}) error {
	if err := modules.NewArticle().Delete(ids); err != nil {
		return err
	}
	return nil
}

func GetArticleById(id interface{}) (*modules.Article, error) {
	if article, err := modules.NewArticle().GetById(id); err != nil {
		return nil, err
	} else {
		return article, nil
	}
}

func PubOrUnpubArticle(a modules.Article) error {
	if err := a.PubOrUnpubById(); err != nil {
		return err
	}
	return nil
}
