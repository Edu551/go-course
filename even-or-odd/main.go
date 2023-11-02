package main

import "fmt"

func main() {
	evenNumbers := []int{}
	oddNumbers := []int{}

	numbers := createNumbers(20)

	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		} else {
			oddNumbers = append(oddNumbers, number)
		}
	}

	fmt.Println("Even ", evenNumbers, "; Odd ", oddNumbers)

}

func createNumbers(r int) []int {
	numbers := []int{}

	for r > 0 {
		numbers = append(numbers, r)
		r -= 1
	}

	return numbers
}
