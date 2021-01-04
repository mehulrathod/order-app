package Routers

import (
	"github.com/gin-gonic/gin"
	helper "updated_structure/orderapp/apiHelpers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Giving access to public folder
	r.Static("/public", "public")

	// for header access
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	uc := hUser()
	cc := hCategory()
	mc := hMenu()
	oc := hOrder()

	v1 := r.Group("/api/v1")
	v1.POST("signup", uc.SignUp)
	v1.POST("login", uc.Login)

	//Authorized routes
	v1.Use(helper.AuthHandler())
	{
		v1.PUT("editProfile", uc.EditProfile)

		v1.POST("addCategory", cc.AddCategory)
		v1.PUT("editCategory", cc.EditCategory)
		v1.GET("getAllCategory", cc.GetAllCategory)
		//v1.POST("statusUpdate", cc.StatusUpdate)

		v1.POST("addMenu", mc.AddMenu)
		v1.PUT("editMenu", mc.EditMenu)
		v1.GET("getAllMenu", mc.GetAllMenu)

		v1.GET("addOrders", oc.AddOrders)
	}

	return r
}