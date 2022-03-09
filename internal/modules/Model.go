package modules

import "time"

type Model struct {
	Id          uint64    `gorm:"primary_key" json:"id"`    // 主键
	CreatedBy   string    `json:"createdBy"`               // 创建人
	UpdateBy    string    `json:"updateBy"`                 // 变更人
	CreatedTime time.Time `json:"createdTime"`              // 创建时间
	UpdateTime  time.Time `json:"updateTime"`               // 变更时间
	Deleted     uint8     `gorm:"column:deleted_" json:"-"` // 删除标志
}
