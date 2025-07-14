package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type CustomClaims struct {
  UserID int64 `json:"userId"`
  jwt.RegisteredClaims
}

func validateJWT(tokenString string) (*CustomClaims, error) {
	tokenString = strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer"))

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		fmt.Println("Auth header:", authHeader)

		if authHeader == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing auth header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := validateJWT(tokenString)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		context.Set("userID", claims.UserID)
		context.Next()
	}
}