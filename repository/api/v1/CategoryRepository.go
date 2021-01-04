package v1

import (
	"github.com/jinzhu/gorm"
	v1res "updated_structure/orderapp/resources/response/api/v1"

	"updated_structure/orderapp/models"
)

type CategoryRepository interface {
	AddCategory(category models.Category)
	EditCategory(category models.Category)
	GetAllCategory()
}

type CategoryRepo struct {
	DB *gorm.DB
}

func (cr *CategoryRepo) GetCategoryById(Id uint) (models.Category, error) {
	category := models.Category{}
	err := cr.DB.Model(category).Select("id, name, status").Where("id = ?", Id).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (cr *CategoryRepo) AddCategory(category models.Category) {
	cr.DB.Model(&category).Create(&category)
	return
}

func (cr *CategoryRepo) EditCategory(category models.Category) {
	cr.DB.Model(&category).Update(&category)
	return
}

func (cr *CategoryRepo) GetAllCategory() ([]v1res.CategoryResponse, error) {
	var categoryList []v1res.CategoryResponse
	cr.DB.Table("categories").Select("categories.*").Where("status = ?", "active").Find(&categoryList)
	return categoryList, nil
}
