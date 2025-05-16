package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return ""
	}
	return string(pass)

}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
