package main

func main() {
	cards := newDeckFromFile("deckOfCards")
	cards.print()
	// cards.saveToFile("deckOfCards")
}
