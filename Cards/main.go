package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, remainingCards := deal(cards, 3)

	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())

	// cards.saveToFile("deckfile")

}
