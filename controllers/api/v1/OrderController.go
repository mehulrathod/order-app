package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "updated_structure/orderapp/apiHelpers"
	v1req "updated_structure/orderapp/resources/request/api/v1"
	v1Service "updated_structure/orderapp/services/api/v1"
)

type OrderController struct {
	OrderService v1Service.OrderService
}

func (oc *OrderController) AddOrders(c *gin.Context)  {
	orderService := oc.OrderService
	var order v1req.OrderRequest
	if err := c.MustBindWith(&order, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := orderService.AddOrder(order)
	helper.Respond(c.Writer, resp)
	return
}