package helpers

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/gin-gonic/gin"
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

func GetTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	bearerToken := strings.Split(authHeader, " ")
	return bearerToken[1], nil
}

func ValidateToken(signedToken string) (*SignedDetails, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok || !token.Valid {
		return nil, errors.New("the token is invalid")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
