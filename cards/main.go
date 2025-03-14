package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Ten of Club")
	for cardNum, card := range cards {
		fmt.Printf("card number %d is %s \n", cardNum+1, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
