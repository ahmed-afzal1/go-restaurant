package models

import "time"

type Cuisine struct {
	ID        uint       `gorm:"primaryKey"`
	Name      *string    `json:"name"`
	Image     *string    `json:"image"`
	Status    uint8      `gorm:"default:1" json:"status"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoCreateTime" json:"updated_at"`

	Restaurants []Restaurant `json:"restaurants" gorm:"foreignKey:CuisineID"`
}
