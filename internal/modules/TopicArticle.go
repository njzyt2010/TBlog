package modules

import "TBlog/pkg/database"

type TopicArticle struct {
	Id        uint64 `gorm:"column:id" json:"id"`
	TopicId   uint64 `json:"topicId" gorm:"column:topic_id"`
	ArticleId uint64 `gorm:"article_id" json:"articleId"`
}

func (t TopicArticle) TableName() string {
	return "t_topic_article"
}

func NewTopicArticle() TopicArticle {
	return TopicArticle{}
}

func (t TopicArticle) Insert() error {
	if err := database.DB.Create(&t).Error; err != nil {
		return err
	}
	return nil
}
func (t TopicArticle) DeleteByTopicId(topicId interface{}) error {
	if err := database.DB.Where("topic_id =?", topicId).Delete(&t).Error; err != nil {
		return err
	}
	return nil
}
func (t TopicArticle) DeleteByArticleId(articleId interface{}) error {
	if err := database.DB.Delete(&t).Where("article_id in (?)", articleId).Error; err != nil {
		return err
	}
	return nil
}
