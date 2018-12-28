package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
//which is a slice of strings
type deck []string

//Create a list of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Receiver function to print
//any variable of type deck now gets access to this method
//sets up methods on variables we've created
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Creating a deal
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Helper function: byte to string transform
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Creating a new file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Reading from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option 1 - log error and return a call to newDeck()
		//Option 2 - log error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
