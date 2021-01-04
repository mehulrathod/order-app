package v1

import (
	"github.com/jinzhu/gorm"
	"updated_structure/orderapp/models"
)

type OrderRepo struct {
	DB *gorm.DB
}

type OrderRepository interface {
	AddOrder(order models.Order)
}

func (or *OrderRepo) AddOrder(order models.Order) {
	or.DB.Model(&order).Create(&order)
}