package auth

import (
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

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Validate Bearer format
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		return nil, jwt.ErrInvalidKey
	}

	// Validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func ValidateTokenAndGetUserID(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims["userId"].(string), nil
}
