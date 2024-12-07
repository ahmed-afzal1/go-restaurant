package models

import (
	"time"
)

type Category struct {
	ID        uint       `gorm:"primaryKey"`
	Name      *string    `json:"name"`
	Slug      *string    `json:"slug"`
	Image     *string    `json:"image"`
	Status    *uint8     `gorm:"default:1" json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	SubCategories []SubCategory `json:"subcategories" gorm:"foreignKey:CategoryID"`
	Food          []Food        `json:"foods" gorm:"foreignKey:CategoryID"`
}
