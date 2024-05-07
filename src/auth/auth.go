package auth

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(userID uint) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 3).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidateTokenAndGetUserID(tokenString string) (uint64, error) {
	// Validate Bearer format
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		return 0, jwt.ErrInvalidKey
	}

	// Validate token
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := permissions["userId"]
		if userIDFloat, ok := userID.(float64); ok {
			return uint64(userIDFloat), nil
		}
	}

	return 0, jwt.ErrInvalidKey
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}
	return config.SecretKey, nil
}
