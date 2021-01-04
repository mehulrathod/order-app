package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	helper "updated_structure/orderapp/apiHelpers"
	v1req "updated_structure/orderapp/resources/request/api/v1"
	v1Service "updated_structure/orderapp/services/api/v1"
)

type MenuController struct {
	MenuService v1Service.MenuService
}

func (mc *MenuController) AddMenu(c *gin.Context) {
	menuService := mc.MenuService
	var menu v1req.AddEditMenuRequest
	if err := c.MustBindWith(&menu, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}

	ImageName, err := helper.ImageUpload(c, "public/images/menus", "menu")
	if err {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Upload valid file."))
		return
	}

	menu.ImageName = ImageName
	resp := menuService.AddMenu(menu)
	helper.Respond(c.Writer, resp)
	return
}

func (mc *MenuController) EditMenu(c *gin.Context) {
	menuService := mc.MenuService
	var menu v1req.AddEditMenuRequest
	if err := c.MustBindWith(&menu, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}

	ImageName, err := helper.ImageUpload(c, "public/images/menus", "menu")
	if err {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Upload valid file."))
		return
	}

	menu.ImageName = ImageName
	resp := menuService.EditMenu(menu)
	helper.Respond(c.Writer, resp)
	return
}

func (mc *MenuController) GetAllMenu(c *gin.Context) {
	resp := mc.MenuService.GetAllMenu()
	helper.Respond(c.Writer, resp)
	return
}
