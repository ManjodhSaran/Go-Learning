package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

func GenerateToken(userID int64, email string) string {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // expire in a week
	})

	token, err := jwtToken.SignedString(secretKey)

	if err != nil {
		return ""
	}

	return token
}

func VerifyToken(token string) (int64, string, error) {
	if len(token) < 7 {
		return 0, "", errors.New("invalid token")
	}
	token = token[7:]

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, "", err
	}
	if !parsedToken.Valid {
		return 0, "", errors.New("invalid token")
	}
	claims := parsedToken.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	email := claims["email"].(string)
	return int64(userID), email, nil

}
