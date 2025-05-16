package models

import (
	"errors"

	"eventapi.com/db"
	"eventapi.com/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `INSERT INTO users(email,password) VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hash := utils.HashPassword(user.Password)

	result, err := stmt.Exec(user.Email, hash)

	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()

	user.ID = userID
	return err
}

func (user User) Validate() (*string, error) {
	query := `SELECT id,password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var password string
	var userID int64

	err := row.Scan(&userID, &password)

	if err != nil {
		return nil, errors.New("User not found")
	}

	if !utils.ComparePassword(user.Password, password) {
		return nil, errors.New("invalid password")
	}

	token := utils.GenerateToken(user.ID, user.Email)
	return &token, nil

}
