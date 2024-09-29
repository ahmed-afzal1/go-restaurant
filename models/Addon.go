package models

import "time"

type Addon struct {
	ID           uint       `gorm:"primaryKey"`
	RestaurantID uint       `json:"restaurant_id"`
	Name         string     `json:"name"`
	Price        float64    `json:"price"`
	CreatedAt    *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"autoCreateTime" json:"updated_at"`

	Restaurant Restaurant `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	Food       []Food     `json:"foods" gorm:"foreignKey:AddonID; references:ID"`
}
