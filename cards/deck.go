package main

import (
	"fmt"
	"math/rand"
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

func getWd() (string, error) {
	// get current working directory
	wd, err := os.Getwd()
	return wd, err
}

func (d deck) saveToFile(fileName string) error {
	// get current working directory
	wd, err := getWd()
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

func newDeckFromFile(fileName string) deck {
	wd, err := getWd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}
	// get []byte from file
	byteSlice, err := os.ReadFile(wd + "\\output\\" + fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}
	// convert to string -> []string -> deck
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {
	// rand.Seed(time.Now().UnixNano())
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		// swap
		d[i], d[newPos] = d[newPos], d[i]
	}
}
