package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int64, email string) string {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // expire in a week
	})

	token, err := jwtToken.SignedString([]byte("secret"))

	if err != nil {
		return ""
	}

	return token
}

func VerifyToken(token string) (int64, string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return 0, "", err
	}
	if !parsedToken.Valid {
		return 0, "", errors.New("invalid token")
	}
	claims := parsedToken.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(int64)
	email := claims["email"].(string)
	return int64(userID), email, nil

}
