package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userID int64) (string, error) {
  secret := os.Getenv("JWT_SECRET")
  if secret == "" {
    return "", errors.New("JWT_SECRET not set in environment")
  }

  claims := jwt.MapClaims{
    "email":  email,
    "userId": userID,
    "exp":    time.Now().Add(2 * time.Hour).Unix(),
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userID int64, jti string) (string, error) {
	// Grab your refresh‐token secret from the environment
	refreshSecret := os.Getenv("REFRESH_SECRET")
	if refreshSecret == "" {
			return "", fmt.Errorf("REFRESH_SECRET not set")
	}

	// Create the claims: sub=subject (user ID), jti=token ID, exp=expiration
	claims := jwt.MapClaims{
			"sub": userID,
			"jti": jti,
			"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	// Build and sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshSecret))
}