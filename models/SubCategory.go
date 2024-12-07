package models

import "time"

type SubCategory struct {
	ID         uint       `gorm:"primaryKey"`
	CategoryID uint       `json:"category_id"`
	Name       *string    `json:"name"`
	Status     uint8      `gorm:"default:1" json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`

	Category Category `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CategoryID;references:ID"`
	Food     []Food   `json:"foods" gorm:"foreignKey:SubCategoryID"`
}
