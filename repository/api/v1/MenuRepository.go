package v1

import (
	"github.com/jinzhu/gorm"
	"updated_structure/orderapp/models"
	v1res "updated_structure/orderapp/resources/response/api/v1"
)

type MenuRepository interface {
	AddMenu(menu models.Menu)
	EditMenu(menu models.Menu)
	GetAllMenu()
}

type MenuRepo struct {
	DB *gorm.DB
}

func (mr *MenuRepo) GetMenuById(Id uint) (models.Menu, error) {
	menu := models.Menu{}
	err := mr.DB.Model(menu).Select("id, name, status, price, image, category_id").Where("id = ?", Id).First(&menu).Error
	if err != nil {
		return menu, err
	}
	return menu, nil
}

func (mr *MenuRepo) GetMenuByName(name string, id uint) (models.Menu, error) {
	menu := models.Menu{}
	err := mr.DB.Model(menu).Select("id, name, status, price, image, category_id").Where("category_id = ?", id).Where("name = ?", name).First(&menu).Error
	if err != nil {
		return menu, err
	}
	return menu, nil
}

func (mr *MenuRepo) AddMenu(menu models.Menu) {
	mr.DB.Model(&menu).Create(&menu)
	return
}

func (mr *MenuRepo) EditMenu(menu models.Menu) {
	mr.DB.Model(&menu).Update(&menu)
	return
}


func (mr *MenuRepo) GetAllMenu() ([]v1res.MenuResponse, error) {
	var menuList []v1res.MenuResponse
	mr.DB.Table("menus").Select("menus.*").Where("status = ?", "active").Find(&menuList)
	return menuList, nil
}
