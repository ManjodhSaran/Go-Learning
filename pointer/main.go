package main

import "fmt"

func main() {

	age := 20

	addAge(&age)
	fmt.Printf("Age: %d\n", age)
}

func addAge(age *int) {
	*age = *age + 18
}
