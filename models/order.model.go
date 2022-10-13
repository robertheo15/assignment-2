package models

import (
	"time"
)

// Order represents the model for an Order
type Order struct {
	ID           uint       `gorm:"primary_key" json:"id" example:"1"`
	CustomerName string     `gorm:"not null;type:varchar(256)" json:"customer_name" binding:"required" example:"Robert"`
	OrderedAt    *time.Time `gorm:"autoCreateTime" json:"ordered_at" example:"2022-11-11T21:21:46+00:00"`
	Items        []Item     `json:"items,omitempty" binding:"required"`
	CreatedAt    *time.Time `json:"created_at,omitempty" example:"2022-11-11T21:21:46+00:00"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" example:"2022-11-11T21:21:46+00:00"`
	DeletedAt    *time.Time `gorm:"index" json:"deleted_at,omitempty" example:"2022-11-11T21:21:46+00:00"`
}
