package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	card2 := "Text"
	card2 = NewCard()

	cards := []string{NewCard(), "Card 2", "Card 1", "Card N"}

	fmt.Println(card)
	fmt.Println(card2)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func NewCard() string {
	return "Card sdfsdf"
}

func Cards(n int) int {
	if n == 0 {
		return 1
	}
	return n + Cards(n-1)
}
