package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := cards.deal(5)
	fmt.Println("Full Deck:")
	cards.print()
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining Cards:")
	remainingCards.print()
}
