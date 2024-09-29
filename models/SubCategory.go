package models

import "time"

type SubCategory struct {
	ID         uint       `gorm:"primaryKey"`
	CategoryID uint       `json:"category_id"`
	Name       *string    `json:"name"`
	Status     uint8      `gorm:"default:1" json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Food     []Food   `json:"foods" gorm:"foreignKey:SubCategoryID"`
}
