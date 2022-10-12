package models

import (
	"gorm.io/gorm"
	"time"
)

// Item represents the model for an item in the item
type Item struct {
	ID          uint           `gorm:"primaryKey" json:"id" example:"1"`
	ItemCode    string         `gorm:"not null;type:varchar(256)" json:"item_code" binding:"required" example:"1"`
	Description string         `gorm:"not null" json:"description" binding:"required" example:"Ini adalah deskripsi"`
	Quantity    int            `gorm:"not null" json:"quantity" binding:"required" example:"1"`
	OrderID     uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order_id,omitempty" example:"1"`
	CreatedAt   *time.Time     `json:"created_at,omitempty" example:"2022-11-11T21:21:46+00:00"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty" example:"2022-11-11T21:21:46+00:00"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" example:"2022-11-11T21:21:46+00:00"`
}
