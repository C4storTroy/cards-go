package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cardsTest := newDeck()
	if len(cardsTest) != 20 {
		t.Errorf("Expected length of 20, but got %v", len(cardsTest))
	}

	if cardsTest[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of spades, but got %v", cardsTest[0])
	}

	if cardsTest[len(cardsTest)-1] != "Five of Clubs" {
		t.Errorf("Expected Five of clubs, but got %v", cardsTest[len(cardsTest)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected length of 20, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
