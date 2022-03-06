package service

import "TBlog/internal/modules"

func InsertMenu(menu modules.Menu) (*modules.Menu, error)  {
	return menu.Insert()
}

func UpdateMenu(menu modules.Menu) error {
	return menu.Update(menu)
}

func DeleteMenu(ids interface{}) error {
	return modules.NewMenu().DeleteByIds(ids)
}