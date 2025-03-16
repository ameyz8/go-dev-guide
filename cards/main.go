package main

func main() {
	// cards := newDeckFromFile("deckOfCards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.saveToFile("deckOfCards")
}
