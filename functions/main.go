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

	fmt.Println("Factorial: ", factorial(3))
	fmt.Println("sumup: ", sumup(3, 10, 7))

}

func sumup(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
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

func factorial(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return n
	}
	return n * factorial(n-1)
}
