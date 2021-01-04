package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"encoding/json"

	helper "updated_structure/orderapp/apiHelpers"
	v1req "updated_structure/orderapp/resources/request/api/v1"
	v1Service "updated_structure/orderapp/services/api/v1"
)

type UserController struct {
	UserService v1Service.UserService
}

func (uc *UserController) SignUp(c *gin.Context) {
	userService := uc.UserService
	var user v1req.SignUpRequest

	if err := c.MustBindWith(&user, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}

	ImageName, err := helper.ImageUpload(c, "public/images/users", "user")
	if err {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Upload valid file."))
		return
	}

	user.ImageName = ImageName

	resp := userService.SignUpUser(user)
	helper.Respond(c.Writer, resp)
	return
}

func (uc *UserController) Login(c *gin.Context) {
	userService := uc.UserService
	var user v1req.LoginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := userService.LoginUser(user)
	helper.Respond(c.Writer, resp)
	return
}

func (uc *UserController) EditProfile(c *gin.Context) {
	userAuth := helper.AuthUser(c)
	UId := uint(userAuth["Id"].(float64))

	userService := uc.UserService
	var user v1req.EditProfileRequest

	if err := c.MustBindWith(&user, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}

	ImageName, err := helper.ImageUpload(c, "public/images/users", "user")
	if err {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Upload valid file."))
		return
	}

	user.ImageName = ImageName
	user.Id = UId

	resp := userService.EditProfile(user)
	helper.Respond(c.Writer, resp)
	return
}