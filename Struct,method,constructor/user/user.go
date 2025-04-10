package user

import (
	"errors"
	"fmt"
	"time"
)

// struct
type User struct {
	name      string
	age       string
	dob       string
	createdAt time.Time
}

// method
func (u *User) Output() {
	fmt.Println("Name: ", u.name)
	fmt.Println("Age: ", u.age)
	fmt.Println("DOB: ", u.dob)
	fmt.Println("CreatedAt: ", u.createdAt)
}

func (u *User) ChangeName(name string) {
	u.name = name
}

type Admin struct {
	email    string
	password string
	user     User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		user: User{
			name: "admin",
			age:  "20",
			dob:  "2020-01-01",
		},
	}
}

func (a Admin) OutputAdmin() {
	fmt.Println("Email: ", a.email)
	fmt.Println("Password: ", a.password)
	a.user.Output()
}

func New(name, age, dob string) (*User, error) {
	if age == "" {
		return nil, errors.New("age must be greater than zero")
	}
	if dob == "" {
		return nil, errors.New("dob is required")
	}
	if name == "" {
		return nil, errors.New("name is required")
	}

	u := &User{
		name:      name,
		age:       age,
		dob:       dob,
		createdAt: time.Now(),
	}

	u.Output()

	return u, nil
}
