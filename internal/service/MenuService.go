package service

import (
	"TBlog/internal/modules"
	"TBlog/internal/repository"
)

var MenuService = newMenuService()

func newMenuService() *menuService {
	return &menuService{}
}

type menuService struct {
}

func (m *menuService) Insert(menu *modules.Menu) error {
	return repository.MenuRepository.Insert(menu)
}

func (m *menuService) Update(menu *modules.Menu) error {
	return repository.MenuRepository.Update(menu)
}
func (m *menuService) Updates(id uint64, values map[string]interface{}) error {
	return repository.MenuRepository.Updates(id, values)
}

func (m *menuService) Delete(id uint64) error {
	return repository.MenuRepository.UpdateColumn(id, "deleted_", true)
}

func (m *menuService) GetMenusOfBlog() []modules.Menu {
	if data, err := repository.MenuRepository.GetMenusOfBlog(); err != nil {
		return []modules.Menu{}
	} else {
		return data
	}
}
