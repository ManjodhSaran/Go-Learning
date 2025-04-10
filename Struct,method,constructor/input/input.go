package input

import "fmt"

func GetUserInput(str string) string {
	fmt.Printf(str, ": ")
	var val string
	fmt.Scanln(&val)
	return val
}
