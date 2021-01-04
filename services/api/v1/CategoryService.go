package v1Service

import (
	"fmt"
	"updated_structure/orderapp/models"

	helpers "updated_structure/orderapp/apiHelpers"
	v1repo "updated_structure/orderapp/repository/api/v1"
	v1req "updated_structure/orderapp/resources/request/api/v1"
	v1res "updated_structure/orderapp/resources/response/api/v1"
)

type CategoryService struct {
	CategoryRequest v1req.CategoryRequest
	Category        models.Category
	CategoryRepo    v1repo.CategoryRepo
}

func (cs *CategoryService) AddCategory(cr v1req.CategoryRequest) map[string]interface{} {
	category := cs.Category
	category.Name = cr.Name
	category.Status = "active"
	cs.CategoryRepo.AddCategory(category)
	userData := v1res.CategoryResponse{
		Name: category.Name,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Category added successfully.")
	response["data"] = userData
	return response
}

func (cs *CategoryService) EditCategory(cr v1req.CategoryRequest) map[string]interface{} {
	category := cs.Category
	category.ID = cr.ID
	category.Name = cr.Name
	category.Status = "active"
	res, err := cs.CategoryRepo.GetCategoryById(category.ID)
	if err != nil {
		fmt.Println(res)
		return helpers.Message(helpers.ResponseError, "Enter valid category.")
	}
	cs.CategoryRepo.EditCategory(category)
	userData := v1res.CategoryResponse{
		Name: category.Name,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Category updated successfully.")
	response["data"] = userData
	return response
}

func (cs *CategoryService) GetAllCategories() map[string]interface{} {
	res, err := cs.CategoryRepo.GetAllCategory()
	if err != nil {
		response := helpers.Message(helpers.ResponseError, "something went wrong.")
		return response
	}
	fmt.Println("all categories", res, err)
	response := helpers.Message(helpers.ResponseSuccess, "success.")
	response["data"] = res
	return response
}