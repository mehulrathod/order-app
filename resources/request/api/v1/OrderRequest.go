package v1

import "time"

type OrderRequest struct {
	ID       uint64    `json:"id"`
	UserId   uint64    `json:"user_id" binding:"required"`
	Quantity int64     `json:"quantity" binding:"required"`
	Status   string    `json:"status"`
	MenuData string    `json:"menu_data" binding:"required"`
	Date     time.Time `form:"date" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type GetOrdersRequest struct {
	ID     uint64
	UserId uint64 `form:"user_id" binding:"required"`
}
