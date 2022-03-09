package modules

import "time"

// Menu 菜单
type Menu struct {
	Model
	Pid    uint64 `gorm:"column:pid_" json:"pid"`
	Name   string `gorm:"column:name_" json:"name"`
	Sorted uint8  `gorm:"column:sorted_" json:"sorted"`
}

func (m Menu) TableName() string {
	return "t_menu"
}

func NewMenu() *Menu {
	curTime := time.Now()
	return &Menu{
		Model: Model{CreatedTime: curTime, UpdateTime: curTime},
	}
}
func UpdateMenu() *Menu {
	curTime := time.Now()
	return &Menu{
		Model: Model{UpdateTime: curTime},
	}
}
