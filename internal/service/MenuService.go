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

func (m *menuService) Delete(ids []uint64) error {
	for i := 0; i < len(ids); i++ {
		if err := repository.MenuRepository.Delete(ids[i]); err != nil {
			return err
		}
	}
	return nil
}

func (m *menuService) GetMenusOfBlog() []modules.Menu {
	if data, err := repository.MenuRepository.GetMenusOfBlog(); err != nil {
		return []modules.Menu{}
	} else {
		return data
	}
}
