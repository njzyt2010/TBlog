package modules

import (
	"TBlog/pkg/database"
	"time"
)

// Menu 菜单
type Menu struct {
	Model
	Pid  uint64 `gorm:"column:pid_" json:"pid"`
	Name string `gorm:"column:name_" json:"name"`
	Sorted uint8 `gorm:"column:sorted_" json:"sorted"`
}

func (m Menu) TableName() string {
	return "t_menu"
}

func NewMenu() Menu {
	return Menu{}
}

// Insert 写入数据
func (m Menu) Insert() (*Menu, error) {
	curTime := time.Now()
	m.CreatedTime = curTime
	m.UpdateTime = curTime

	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (m Menu) Update(values interface{}) error {

	if err := database.DB.Model(&m).Where("id = ?", m.Id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (m Menu) DeleteByIds(ids interface{}) error {
	if err:=database.DB.Exec("UPDATE t_menu SET deleted_ = 1 WHERE id IN(?)" ,ids).Error ; err !=nil {
		return err
	}
	return nil
}

func (m Menu) GetMenusOfBlog() ([]Menu ,error) {
	var menus []Menu
	if err := database.DB.Model(&m).Where("deleted_ = ?", 0).Find(&menus).Error;err !=nil {
		return nil ,err
	}
	return menus,nil

}