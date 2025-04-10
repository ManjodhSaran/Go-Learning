package user

import (
	"errors"
	"fmt"
	"time"
)

// struct
type user struct {
	name      string
	age       int
	dob       string
	createdAt time.Time
}

// method
func (u *user) output() {
	fmt.Println("Name: ", u.name)
	fmt.Println("Age: ", u.age)
	fmt.Println("DOB: ", u.dob)
	fmt.Println("CreatedAt: ", u.createdAt)
}

func (u *user) changeName(name string) {
	u.name = name
}

func User(name string, age int, dob string) (*user, error) {
	if age <= 0 {
		return nil, errors.New("Age must be greater than zero")
	}
	if dob == "" {
		return nil, errors.New("DOB is required")
	}
	if name == "" {
		return nil, errors.New("Name is required")
	}
	u := &user{
		name:      name,
		age:       age,
		dob:       dob,
		createdAt: time.Now(),
	}

	u.output()

	return u, nil
}
