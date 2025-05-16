package models

import (
	"errors"

	"eventapi.com/db"
	"eventapi.com/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserData struct {
	ID    int64  `json:"id"`
	Email string `json:"email" binding:"required"`
}

func GetAllUsers(users *[]UserData) error {
	query := `SELECT id,email FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserData
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return err
		}
		*users = append(*users, user)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (user *User) Save() error {
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

func (user *User) Validate() (*string, error) {
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
	user.ID = userID
	token := utils.GenerateToken(user.ID, user.Email)
	return &token, nil

}
