package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	// get current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	// create new directory
	err = os.MkdirAll(wd+"\\output\\", 0666)
	if err != nil {
		return err
	}
	// save to new directory
	return os.WriteFile(wd+"\\output\\"+fileName, []byte(d.toString()), 0666)
}
