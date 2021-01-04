package models

type Category struct {
	Model
	Name   string `gorm:"type:varchar(60)" json:"name" validate:"required"`
	Status string `gorm:"type:enum('active', 'inactive')" json:"status"`
}

func (u *Category) TableName() string {
	return "categories"
}
