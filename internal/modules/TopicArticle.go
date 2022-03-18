package modules

type TopicArticle struct {
	 TopicId uint64 `gorm:"column:topic_id" json:"topicId"`	//栏目id
	 ArticleId uint64 `gorm:"column:article_id" json:"articleId"` // 文章id
}
