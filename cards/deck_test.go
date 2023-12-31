package main

import (
	"os"
	"testing"
)

func Test_NewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Ace of Spades, but got %v", d[len(d)-1])
	}
}

func Test_Save_To_Deck_And_NewDeck_From_File(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 Cards in Deck, got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
