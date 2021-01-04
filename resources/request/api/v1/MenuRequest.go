package v1

import (
	"mime/multipart"
	"time"
)

type AddEditMenuRequest struct {
	ID         uint                  `form:"id"`
	Name       string                `form:"name" binding:"required"`
	Price      float64               `form:"price" binding:"required"`
	Image      *multipart.FileHeader `form:"image" binding:"required"`
	ImageName  string
	Status     string    `form:"status"`
	CategoryId uint      `form:"category_id" binding:"required"`
	CreatedAt  time.Time `form:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
