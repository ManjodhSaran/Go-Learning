package main

import (
	"SMC/input"
	"SMC/user"
)

func main() {
	name := input.GetUserInput("Enter your name: ")
	age := input.GetUserInput("Enter your age: ")
	dob := input.GetUserInput("Enter your DOB: ")

	new_user, err := user.New(name, age, dob)

	if err != nil {
		panic(err)
	}

	new_user.Output()
	new_user.ChangeName("John")
	new_user.Output()

	admin := user.NewAdmin("admin@gmail.com", "123456")
	admin.OutputAdmin()
	admin.user.Output()
}
