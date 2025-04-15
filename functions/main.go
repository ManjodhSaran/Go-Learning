package main

import "fmt"

type funcType func(int) int

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, func(n int) int {
		return n * 2
	})
	tripled := transformNumbers(&numbers, createTransformer(3))

	fmt.Println("original: ", numbers)
	fmt.Println("doubled: ", doubled)
	fmt.Println("tripled: ", tripled)
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}

func transformNumbers(numbers *[]int, transform funcType) []int {

	newNums := []int{}

	for _, num := range *numbers {
		newNums = append(newNums, transform(num))
	}

	return newNums

}

func getTransformerFunc(n int) funcType {
	if n > 3 {
		return double
	}
	return triple
}

func createTransformer(factor int) funcType {
	return func(n int) int {
		return n * factor
	}
}
