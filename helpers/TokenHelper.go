package helpers

import (
	"log"
	"os"
	"time"

	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Email string
	Uid   uint
	jwt.RegisteredClaims
}

func GenerateToken(email string, uid uint) (signedToken string, signedRefreshToken string, err error) {
	var user models.User

	claims := &SignedDetails{
		Email: email,
		Uid:   uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 168)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	if err := config.DB.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
		"token":         token,
		"refresh_token": refreshToken,
	}).Error; err != nil {
		return token, refreshToken, err
	}

	return token, refreshToken, err
}
