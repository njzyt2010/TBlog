package repository

import (
	"TBlog/internal/modules"
	"TBlog/pkg/database"
)

type menuRepository struct {
}

func newMenuRepository() *menuRepository {
	return &menuRepository{}
}

var MenuRepository = newMenuRepository()

// 新增菜单
func (m *menuRepository) Insert(menu *modules.Menu) error {
	err := database.DB.Create(menu).Error
	return err
}

// 修改菜单
func (m *menuRepository) Update(menu *modules.Menu) error {
	err := database.DB.Save(menu).Error
	return err
}

//修改菜单
func (m *menuRepository) Updates(id uint64, values map[string]interface{}) error {
	err := database.DB.Model(&modules.Menu{}).Where("id = ?", id).Updates(values).Error
	return err
}

func (m *menuRepository) UpdateColumn(id uint64, column string, value interface{}) error {
	err := database.DB.Model(&modules.Menu{}).Where("id = ? ", id).UpdateColumn(column, value).Error
	return err
}

//删除菜单
func (m *menuRepository) Delete(id uint64) {
	database.DB.Delete(&modules.Menu{}, "id = ?", id)
}

// 查询博客下所有的菜单
func (m *menuRepository) GetMenusOfBlog() ([]modules.Menu, error) {
	var menus []modules.Menu = nil
	if err := database.DB.Model(&modules.Menu{}).Where("deleted_ = ?", 0).Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}