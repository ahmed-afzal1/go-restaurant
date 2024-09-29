package models

import "time"

type Restaurant struct {
	ID              uint       `gorm:"primaryKey"`
	CuisineID       uint       `json:"cuisine_id"`
	Name            *string    `json:"name"`
	Address         *string    `json:"address"`
	Logo            *string    `json:"logo"`
	VatTax          *uint16    `json:"vat_tax"`
	MinDeliveryTime *string    `json:"min_delivery_time"`
	Latitude        *string    `json:"latitude"`
	Longtitude      *string    `json:"longtitude"`
	OwnerFirstName  *string    `json:"owner_first_name"`
	OwnerLastName   *string    `json:"owner_last_name"`
	OwnerEmail      *string    `json:"owner_email"`
	Password        *string    `json:"password"`
	ConfirmPassword *string    `json:"confirm_password"`
	OwnerPhoneNo    *string    `json:"owner_phone_no"`
	Tin             *string    `json:"tin"`
	LicenseDocument *string    `json:"license_document"`
	Tags            *string    `json:"tags"`
	CreatedAt       *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"autoCreateTime" json:"updated_at"`

	Cuisine Cuisine `json:"cuisine" gorm:"foreignKey:CuisineID"`
	Addon   []Addon `json:"addons" gorm:"foreignKey:RestaurantID"`
	Food    []Food  `json:"foods" gorm:"foreignKey:RestaurantID"`
}
