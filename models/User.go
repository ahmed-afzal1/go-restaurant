package models

import (
	"time"
)

type User struct {
	ID            uint       `gorm:"primaryKey"`
	First_name    *string    `json:"first_name"`
	Last_name     *string    `json:"last_name"`
	Email         *string    `json:"email"`
	Phone         *string    `json:"phone"`
	Password      *string    `json:"password"`
	Token         *string    `json:"token"`
	Refresh_token *string    `json:"refresh_token"`
	Created_at    *time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at    *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
