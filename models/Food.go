package models

import "time"

type Food struct {
	ID               uint       `gorm:"primaryKey"`
	RestaurantID     uint       `json:"restaurant_id"`
	CategoryID       uint       `json:"category_id"`
	SubCategoryID    uint       `json:"subcategory_id"`
	AddonID          uint       `json:"addon_id"`
	Name             *string    `json:"name"`
	Slug             *string    `json:"slug"`
	Description      *string    `json:"description"`
	Image            *string    `json:"image"`
	IsHalal          uint8      `gorm:"default:0" json:"is_halal"`
	FoodType         *string    `json:"food_type"`
	AddonStartTime   *time.Time `json:"addon_start_time"`
	AddonEndTime     *time.Time `json:"addon_end_time"`
	Price            float32    `json:"price"`
	DiscountType     *string    `json:"discount_type"`
	Discount         *float32   `json:"discount"`
	MaxPurchaseLimit *uint16    `json:"max_purchase_limit"`
	Tags             *string    `json:"tags"`

	Restaurant  Restaurant  `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	Category    Category    `json:"category" gorm:"foreignKey:CategoryID"`
	SubCategory SubCategory `json:"subcategory" gorm:"foreignKey:SubCategoryID"`
	Addon       Addon       `json:"addon" gorm:"foreignKey:AddonID"`
}
