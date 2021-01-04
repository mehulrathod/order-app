package models

type Order struct {
	Model
	UserId   uint64   `json:"user_id" validate:"required"`
	Status   string `gorm:"type:enum('active', 'inactive')" json:"status"`
	MenuData string `json:"menu_data" binding:"required"`
	Quantity int64  `gorm:"not mull" json:"quantity"`
}

func (u *Order) TableName() string {
	return "order"
}
