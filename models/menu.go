package models

type Menu struct {
	Model
	Name       string  `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
	Image      string  `gorm:"type:varchar(60)" json:"image"`
	Status     string  `gorm:"type:enum('active', 'inactive')" json:"status"`
	CategoryId uint    `json:"category_id" validate:"required"`
}

func (u *Menu) TableName() string {
	return "menus"
}
