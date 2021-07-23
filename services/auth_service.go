package services

import (
	"fmt"
	"lander/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint, email string) (string, error) {
	config, err := configs.LoadConfig(".")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := jwtToken.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	config, _ := configs.LoadConfig(".")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(config.JWT_SECRET), nil
	})

}
