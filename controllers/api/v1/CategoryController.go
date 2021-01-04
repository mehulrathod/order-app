package v1

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	helper "updated_structure/orderapp/apiHelpers"
	v1req "updated_structure/orderapp/resources/request/api/v1"
	v1Service "updated_structure/orderapp/services/api/v1"
)

type CategoryController struct {
	CategoryService v1Service.CategoryService
}

func (cc *CategoryController) AddCategory(c *gin.Context) {
	categoryService := cc.CategoryService
	var category v1req.CategoryRequest
	err := json.NewDecoder(c.Request.Body).Decode(&category)
	if err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := categoryService.AddCategory(category)
	helper.Respond(c.Writer, resp)
	return
}

func (cc *CategoryController) EditCategory(c *gin.Context) {
	categoryService := cc.CategoryService
	var category v1req.CategoryRequest
	err := json.NewDecoder(c.Request.Body).Decode(&category)
	if err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := categoryService.EditCategory(category)
	helper.Respond(c.Writer, resp)
	return
}

func (cc *CategoryController) GetAllCategory(c *gin.Context) {
	resp := cc.CategoryService.GetAllCategories()
	helper.Respond(c.Writer, resp)
	return
}

/*func (cc *CategoryController) StatusUpdate(c *gin.Context) {
	var category v1req.StatusUpdateRequest
	resp := cc.CategoryService.StatusUpdate()
	helper.Respond(c.Writer, resp)
	return
}
*/