package main

import (
	"os"
	"testing"
)

//With capital "T"

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expected 52 card but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//To remove any previous test files
	os.Remove("_deckTesting")

	deck := newDeck()               //New Deck
	deck.saveToFile("_deckTesting") //Save to file - Reciver function

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52, got %v", len(loadedDeck))
	}

	//clean up the file after the test
	os.Remove("_deckTesting")
}
