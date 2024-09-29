package services

import (
	"errors"

	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
	"github.com/ahmed-afzal1/restaurant/requests"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Signup(req requests.RegisterRequest) (models.User, error) {
	var user models.User
	if req.Password != req.PasswordConfirmation {
		return models.User{}, errors.New("passwords do not match")
	}

	password, _ := HashPassword(req.Password)
	req.Password = password
	req.PasswordConfirmation = password

	user, err := repositories.Signup(req)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
