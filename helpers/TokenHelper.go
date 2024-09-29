package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

var secret_key = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Email string
	Uid   uint
	jwt.RegisteredClaims
}

func GenerateToken(email string, uid uint) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Uid:   uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(secret_key))

	if err != nil {
		log.Panic(err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
