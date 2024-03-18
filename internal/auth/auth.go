package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "a9a83936-5e17-4f99-8cee-854eba97585f"

type Account struct {
	Username string
	Role     int
}

func GenerateToken(acc Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": acc.Username,
		"role":     acc.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(signedToken string) (*Account, error) {
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		username := claims["username"].(string)
		role := int(claims["role"].(float64))
		return &Account{
			Username: username,
			Role:     role,
		}, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
