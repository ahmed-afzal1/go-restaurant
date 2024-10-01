package repositories

import (
	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/helpers"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/requests"
)

func Signup(req requests.RegisterRequest) (models.User, error) {
	var user models.User

	user.Email = &req.Email
	user.Phone = &req.Phone
	user.Password = &req.Password

	result := config.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	token, refreshToken, _ := helpers.GenerateToken(*user.Email, user.ID)

	user.Token = &token
	user.Refresh_token = &refreshToken

	return user, nil
}

func Login(req requests.LoginRequest) (models.User, error) {
	var user models.User

	result := config.DB.Where("email =?", req.Email).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	token, refreshToken, _ := helpers.GenerateToken(*user.Email, user.ID)

	user.Token = &token
	user.Refresh_token = &refreshToken

	return user, nil
}
