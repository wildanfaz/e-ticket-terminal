package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wildanfaz/e-ticket-terminal/configs"
)

type NewClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	cfg := configs.InitConfig()

	claims := NewClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWTDuration)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString(cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (*NewClaims, error) {
	cfg := configs.InitConfig()

	jwtToken, err := jwt.ParseWithClaims(token, &NewClaims{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*NewClaims)
	if !ok {
		return nil, errors.New("Invalid Claims")
	}

	return claims, nil
}
