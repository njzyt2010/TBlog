package modules

import "time"

type Model struct {
	Id          uint    `gorm:"primary_key" json:"id"`
	CreatedBy   string    ` json:"createdBy"`
	UpdateBy    string    `json:"updateBy"`
	CreatedTime time.Time `json:"createdTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Deleted     uint8     `gorm:"column:deleted_" json:"-"`
}
