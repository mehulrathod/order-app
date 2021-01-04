package v1

import (
	"time"
)

type CategoryRequest struct {
	ID        uint
	Name      string    `form:"name" binding:"required"`
	Status    string    `form:"status" binding:"required"`
	CreatedAt time.Time `form:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type StatusUpdateRequest struct {
	ID        uint
	Status    string    `form:"status" binding:"required"`
}
