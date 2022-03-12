package modules

type ArticleTag struct {
	Id        uint64 `gorm:"column:id" json:"id"`                //主键
	ArticleId uint64 `gorm:"column:article_id" json:"articleId"` //文章id
	TagId     uint64 `gorm:"column:tag_id" json:"tagId"`         // 标签id
}

func (a *ArticleTag) TableName() string {
	return "t_article_tag"
}
