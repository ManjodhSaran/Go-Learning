package main

import "fmt"

func main() {
	hobbies := []string{"reading", "gaming", "coding"}

	fmt.Println("My hobbies are:")
	for i, hobby := range hobbies {
		fmt.Printf("%d. %s\n", i+1, hobby)
	}

}
