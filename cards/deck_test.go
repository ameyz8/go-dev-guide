package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Length of Deck is not equal to 16, instead is %v", len(d))
	}

	// Assertion of expectation for first card
	if d[0] != "Ace of Spades" {
		t.Errorf("Expcted first card to be Ace of Spades, instead it is %v", d[0])
	}

	// Assertion of expectation for last card
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expcted last card to be Four of Clubs, instead it is %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	wd, _ := getWd()
	os.Remove(wd + "\\output\\" + "_deckTesting")

	deck := newDeck()

	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected the length of the deck to be 16, but instead it is %v", len(deck))
	}

	os.Remove(wd + "\\output\\" + "_deckTesting")
}
