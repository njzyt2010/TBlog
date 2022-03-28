package modules

import "time"

// 文章
type Article struct {
	Model
	Title      string   `gorm:"column:title_" json:"title"`           //文章标题
	Content    string   `gorm:"column:content_" json:"content"`       //内容
	Reprint    *bool    `gorm:"column:reprint_" json:"reprint"`       // 是否为转载的文章
	ReprintUrl string   `gorm:"column:reprint_url" json:"reprintUrl"` // 转载的地址
	Published  *bool    `gorm:"column:published_" json:"published"`   //文章是否发布
	PublishedTime time.Time `gorm:"column:published_time" json:"publishedTime"` // 文章发布时间
	TopicId    uint64   `gorm:"column:topic_id" json:"topicId"`        //文章所属栏目
	TagId      []uint64 `gorm:"-"`                                    //文章标签
}

func (a Article) TableName() string {
	return "t_article"
}

func NewArticle() *Article {
	var curTime = time.Now()
	return &Article{
		Model: Model{
			CreatedTime: curTime,
			UpdateTime:  curTime,
			Deleted:     false,
		},
	}
}
func UpdateArticle() *Article {
	return &Article{
		Model: Model{UpdateTime: time.Now()},
	}
}
