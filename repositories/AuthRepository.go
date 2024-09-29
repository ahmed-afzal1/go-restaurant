package repositories

import (
	"go/token"

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

	if err := config.DB.Create(&user); err != nil {
		return models.User{}, err.Error
	}

	token, refreshToken, _ : = helpers.GenerateToken(*user.Email, user.ID)
	user.Token = &token
	user.Refresh_token = &refreshToken

	return user, nil
}
