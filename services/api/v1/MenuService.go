package v1Service

import (
	"fmt"
	"updated_structure/orderapp/models"

	helpers "updated_structure/orderapp/apiHelpers"
	v1repo "updated_structure/orderapp/repository/api/v1"
	v1req "updated_structure/orderapp/resources/request/api/v1"
	v1res "updated_structure/orderapp/resources/response/api/v1"
)

type MenuService struct {
	MenuRequest v1req.AddEditMenuRequest
	Menu        models.Menu
	MenuRepo    v1repo.MenuRepo
}

func (ms *MenuService) AddMenu(mr v1req.AddEditMenuRequest) map[string]interface{} {
	menu := ms.Menu
	menu.Name = mr.Name
	menu.Price = mr.Price
	menu.Image = mr.ImageName
	menu.Status = "active"
	menu.CategoryId = mr.CategoryId

	//check menu is already exist or not
	res, err := ms.MenuRepo.GetMenuByName(menu.Name, menu.CategoryId)
	if err == nil {
		fmt.Println(res)
		helpers.ImageDelete("public/images/menus/" + menu.Image)
		return helpers.Message(helpers.ResponseError, "Menu name is already exists with this category.")
	}

	ms.MenuRepo.AddMenu(menu)
	userData := v1res.MenuResponse{
		Name:  menu.Name,
		Price: menu.Price,
		Image: "http://localhost:" + helpers.GetPort() + "/public/images/menus/" + menu.Image,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Menu added successfully.")
	response["data"] = userData
	return response
}

func (ms *MenuService) EditMenu(mr v1req.AddEditMenuRequest) map[string]interface{} {
	menu := ms.Menu
	menu.ID = mr.ID
	menu.Name = mr.Name
	menu.Price = mr.Price
	menu.Image = mr.ImageName
	menu.Status = mr.Status
	menu.CategoryId = mr.CategoryId

	fmt.Println("service ====>", menu.ID)

	//check menu is exist or not
	menuId, err := ms.MenuRepo.GetMenuById(menu.ID)
	if err != nil {
		fmt.Println(menuId)
		helpers.ImageDelete("public/images/menus/" + menu.Image)
		return helpers.Message(helpers.ResponseError, "Select valid menu.")
	}

	//check menu is already exist or not
	res, err := ms.MenuRepo.GetMenuByName(menu.Name, menu.CategoryId)
	if err == nil {
		fmt.Println(res)
		helpers.ImageDelete("public/images/menus/" + menu.Image)
		return helpers.Message(helpers.ResponseError, "Menu name is already exists with this category.")
	}

	ms.MenuRepo.EditMenu(menu)
	helpers.ImageDelete("public/images/users/" + res.Image)
	userData := v1res.MenuResponse{
		Name:  menu.Name,
		Price: menu.Price,
		Image: "http://localhost:" + helpers.GetPort() + "/public/images/menus/" + menu.Image,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Menu updated successfully.")
	response["data"] = userData
	return response
}

func (ms *MenuService) GetAllMenu() map[string]interface{} {
	res, err := ms.MenuRepo.GetAllMenu()
	if err != nil {
		response := helpers.Message(helpers.ResponseError, "something went wrong.")
		return response
	}
	response := helpers.Message(helpers.ResponseSuccess, "GetAllMenu successfully.")
	response["data"] = res
	return response
}
