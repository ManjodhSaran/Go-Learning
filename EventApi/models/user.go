package models

import (
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
