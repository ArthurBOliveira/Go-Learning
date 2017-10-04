package main

func main() {
	cards := deck{newCard(), "Ace of Spades"}
	cards = append(cards, "King of Hearts")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
