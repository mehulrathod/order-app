package models

type User struct {
	Model
	Name     string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(60)" json:"email" validate:"required,email"`
	Image    string `gorm:"type:varchar(60)" json:"image"`
	Password string `gorm:"type:varchar(60)" json:"password" validate:"required"`
	Mobile   int64  `gorm:"type:int(10)" json:"mobile" validate:"required"`
}

func (u *User) TableName() string {
	return "users"
}
