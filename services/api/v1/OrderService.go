package v1Service

import (
	helpers "updated_structure/orderapp/apiHelpers"
	"updated_structure/orderapp/models"
	v1repo "updated_structure/orderapp/repository/api/v1"
	v1req "updated_structure/orderapp/resources/request/api/v1"
	v1res "updated_structure/orderapp/resources/response/api/v1"
)

type OrderService struct {
	OrderRequest v1req.OrderRequest
	Order        models.Order
	OrderRepo    v1repo.OrderRepo
}


func (os *OrderService) AddOrder(or v1req.OrderRequest) map[string]interface{} {
	order := os.Order
	order.MenuData = or.MenuData
	order.Status = or.Status
	order.UserId = or.UserId
	os.OrderRepo.AddOrder(order)
	userData := v1res.OrderResponse{
		MenuData: order.MenuData,
		Status: order.Status,
		UserId: order.UserId,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Order added successfully.")
	response["data"] = userData
	return response
}