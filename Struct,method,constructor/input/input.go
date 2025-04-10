package input

import "fmt"

func GetUserInput(str string) string {
	fmt.Print(str)
	var val string
	fmt.Scanln(&val)
	return val
}
