package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := []int{}
	oddNumbers := []int{}

	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		} else {
			oddNumbers = append(oddNumbers, number)
		}
	}

	fmt.Println("Even ", evenNumbers, "; Odd ", oddNumbers)

}
