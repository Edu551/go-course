package main

func main() {
	cards := newDeckFromFile("my_cards")

	cards.print()
}

func Cards(n int) int {
	if n == 0 {
		return 1
	}
	return n + Cards(n-1)
}
