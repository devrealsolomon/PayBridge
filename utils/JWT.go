package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(payload map[string]interface{}, options jwt.MapClaims) (string, error) {
	privateKey := []byte(os.Getenv("JWT_SECRETS"))
	defaultOptions := jwt.MapClaims{
		"exp": time.Now().Add(3 * time.Hour).Unix(),
	}

	for key, value := range options {
		defaultOptions[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, defaultOptions)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	publicKey := []byte(os.Getenv("JWT_SECRETS"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
