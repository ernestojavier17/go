package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	)


// Create a new type of 'deck' which is a slice of strings
type deck []string //'deck === []sting' Comparing to OO: this means kind of extend or borrow all the behaviour of a slice string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Tree"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

//This is a receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//Option #1 - Log the error and return a call to newDeck()
		//Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {
	//source := rand.NewSource()
	//r := rand.New(source)

	for i := range d {
		r := rand.Intn(len(d) - 1)
		fmt.Print(" ",r)
		newPos := r

		d[i], d[newPos] = d[newPos], d[i]
	}
}

