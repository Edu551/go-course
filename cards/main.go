package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}

func Cards(n int) int {
	if n == 0 {
		return 1
	}
	return n + Cards(n-1)
}
